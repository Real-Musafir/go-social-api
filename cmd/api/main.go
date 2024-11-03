package main

import (
	"log"

	"github.com/Real-Musafir/social/internal/db"
	env "github.com/Real-Musafir/social/internal/env"
	"github.com/Real-Musafir/social/internal/store"
)

func main() {

	cfg := config {
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://shahadath@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConnes: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConnes, cfg.db.maxIdleTime)
	if err!=nil{
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")


	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	

	mux := app.mount()

	log.Fatal(app.run(mux))
}