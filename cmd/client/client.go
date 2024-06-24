package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
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

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			m := Message{UserID: generateRandomInt(5, 1), Message: getRandomMessage()}
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
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
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
