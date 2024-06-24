package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/karthikraobr/bot/internal/infra/server"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

func main() {
	flag.Parse()
	server := server.NewServer()
	http.Handle("/", server.Router())
	log.Printf("server started at %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
