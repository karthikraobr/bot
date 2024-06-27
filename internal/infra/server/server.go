package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/karthikraobr/bot/internal/infra/store"
	psql "github.com/karthikraobr/bot/internal/querier"
	"github.com/karthikraobr/bot/internal/reponders"
)

var re = regexp.MustCompile(`\d+`)

type server struct {
	upgrade      *websocket.Upgrader
	logger       *log.Logger
	responders   map[string]Responder
	messageCache map[int][]reponders.Message
	store        Store
}

type Responder interface {
	Answer(ctx context.Context, message reponders.Message, messageContext []reponders.Message) string
	Type() string
}

type Store interface {
	InsertReview(ctx context.Context, arg psql.InsertReviewParams) (int32, error)
}

func NewServer(ctx context.Context, pool *pgxpool.Pool) *server {
	return &server{
		upgrade:      &websocket.Upgrader{},
		logger:       log.New(os.Stdout, "server: ", log.LstdFlags),
		messageCache: make(map[int][]reponders.Message),
		responders: map[string]Responder{
			reponders.TypeReview:  &reponders.ReviewResponder{},
			reponders.TypeSupport: &reponders.SupportResponder{},
			reponders.TypeDefault: &reponders.DefaultResponder{},
		},
		store: store.NewStore(ctx, pool),
	}
}

func (s *server) Router() *http.ServeMux {
	mux := http.ServeMux{}
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/", s.wsHandler)
	return &mux
}

type Message struct {
	UserID  int
	Message string
}

func (s *server) wsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c, err := s.upgrade.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Print("upgrade:", err)
		return
	}
	c.SetCloseHandler(s.Closer(ctx))
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			s.logger.Println("read:", err)
			break
		}
		var m Message
		if err := json.Unmarshal(message, &m); err != nil {
			s.logger.Println("write:", err)
			break
		}
		responder := s.getResponder()

		msg := reponders.Message{
			Message: m.Message,
			Time:    time.Now().UTC(),
		}

		s.logger.Println(m)

		s.messageCache[m.UserID] = append(s.messageCache[m.UserID], msg)

		answer := responder.Answer(ctx, msg, s.messageCache[m.UserID])

		err = c.WriteMessage(mt, []byte(answer))
		if err != nil {
			s.logger.Println("write:", err)
			break
		}
	}
}

func (s *server) getResponder() Responder {
	return s.responders[reponders.TypeReview]
}

func (s *server) Closer(ctx context.Context) func(code int, text string) error {
	return func(code int, text string) error {
		id, err := strconv.Atoi(text)
		if err != nil {
			return err
		}
		if messages, ok := s.messageCache[id]; ok {
			msg := messages[len(messages)-1].Message
			r := re.FindString(msg)
			rating, err := strconv.Atoi(r)
			if err != nil {
				return err
			}
			if _, err := s.store.InsertReview(context.Background(), psql.InsertReviewParams{
				Rating:         int32(rating),
				ProductID:      pgtype.Int4{Int32: int32(1), Valid: true},
				CustomerID:     pgtype.Int4{Int32: int32(id), Valid: true},
				ReviewText:     pgtype.Text{String: msg, Valid: true},
				CreatedAt:      pgtype.Timestamp{Time: time.Now(), Valid: true},
				LastModifiedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
			}); err != nil {
				return err
			}
		}
		return nil
	}
}
