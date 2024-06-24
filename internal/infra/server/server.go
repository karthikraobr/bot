package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/karthikraobr/bot/internal/reponders"
)

type server struct {
	upgrade      *websocket.Upgrader
	logger       *log.Logger
	responders   map[string]Responder
	messageCache map[int][]reponders.Message
}

type Responder interface {
	Answer(ctx context.Context, message reponders.Message, messageContext []reponders.Message) string
	Type() string
}

func NewServer() *server {
	return &server{
		upgrade:      &websocket.Upgrader{},
		logger:       log.New(os.Stdout, "server: ", log.LstdFlags),
		messageCache: make(map[int][]reponders.Message),
		responders: map[string]Responder{
			reponders.TypeReview:  &reponders.ReviewResponder{},
			reponders.TypeSupport: &reponders.SupportResponder{},
			reponders.TypeDefault: &reponders.DefaultResponder{},
		},
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
	defer c.Close()
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
		responder := s.getResponder(string(m.Message))

		msg := reponders.Message{
			Message: string(message),
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

func (s *server) getResponder(message string) Responder {
	if strings.Contains(message, "review") {
		return s.responders[reponders.TypeReview]
	} else if strings.Contains(message, "support") {
		return s.responders[reponders.TypeSupport]
	}
	return s.responders[reponders.TypeDefault]
}
