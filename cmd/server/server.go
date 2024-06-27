package main

import (
	"context"
	"log"
	"net/http"

	"github.com/karthikraobr/bot/internal/db"
	"github.com/karthikraobr/bot/internal/infra/server"
)

func main() {
	ctx := context.Background()
	c, err := Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	db, err := db.NewDB(ctx, c.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	server := server.NewServer(ctx, db)
	http.Handle("/", server.Router())
	log.Printf("server started at %s", c.Address)
	log.Fatal(http.ListenAndServe(c.Address, nil))
}
