package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

type Message struct {
	UserID  int
	Message string
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	msg := make(chan string)

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			msg <- string(message)
			log.Printf("recv: %s", message)
		}
	}()

	reponses := []string{"Hi", "Sure, I can do that.", "I'd give it a 5.", ""}

	s := &sync.Once{}
	s.Do(func() {
		m := Message{UserID: 1, Message: "Hi"}
		b, err := json.Marshal(m)
		if err != nil {
			log.Println("marshal:", err)
			return
		}
		err = c.WriteMessage(websocket.TextMessage, b)
		if err != nil {
			log.Println("write:", err)
			return
		}
	})
	for {
		select {
		case <-done:
			return
		case m := <-msg:
			if strings.Contains(m, "Thank you") {
				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "1"))
				if err != nil {
					log.Println("write close:", err)
					return
				}
			}
			finalMessage := Message{UserID: 1}
			if strings.Contains(m, "share your thoughts") {
				finalMessage.Message = reponses[1]
			} else {
				finalMessage.Message = reponses[2]
			}
			b, err := json.Marshal(finalMessage)
			if err != nil {
				log.Println("marshal:", err)
				return
			}
			err = c.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}

// GenerateRandomInt generates a random integer within the specified range [min, max].
func generateRandomInt(max, min int) int {
	// Seed the random number generator to ensure different results each time
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate and return a random integer within the specified range
	return rand.Intn(max-min+1) + min
}

func getRandomMessage() string {
	messages := []string{"review", "support", "default"}
	return messages[generateRandomInt(len(messages)-1, 0)]
}
