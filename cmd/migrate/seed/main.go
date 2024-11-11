package main

import (
	"log"

	"github.com/Real-Musafir/social/internal/db"
	"github.com/Real-Musafir/social/internal/store"
)

func main() {
	addr := "postgres://shahadath@localhost/socialnetwork?sslmode=disable"
	
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}