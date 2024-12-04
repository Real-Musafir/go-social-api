package main

import (
	"github.com/Real-Musafir/social/internal/db"
	env "github.com/Real-Musafir/social/internal/env"
	"github.com/Real-Musafir/social/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

func main() {

	cfg := config {
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://shahadath@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConnes: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConnes, cfg.db.maxIdleTime)
	if err!=nil{
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("Database connection pool established")


	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	

	mux := app.mount()

	logger.Fatal(app.run(mux))
}