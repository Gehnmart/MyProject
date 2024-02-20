package main

import (
	"MyProject/EasyMessage/server/internal/model"
	"MyProject/EasyMessage/server/internal/user"
	"MyProject/EasyMessage/server/internal/ws"
	"MyProject/EasyMessage/server/router"
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const verson = "1.0.0"

type config struct {
	addres string
	port   string
	env    string
	db     struct {
		dsn             string
		maxOpenConns    int
		maxIdleConns    int
		connMaxIdleTime string
	}
}

type application struct {
	config   config
	logger   *log.Logger
	models   *model.Models
	handlers *user.Handler
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addres, "address", "localhost", "API server address")
	flag.StringVar(&cfg.port, "port", "4000", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dns", os.Getenv("EASYMESSAGE_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.connMaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max idle time")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	userModel := model.NewModel(db)
	userHandler := user.NewHandler(&userModel.Users)

	app := application{
		config:   cfg,
		logger:   logger,
		models:   userModel,
		handlers: userHandler,
	}

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(app.handlers, wsHandler)
	router.Start(cfg.addres + ":" + cfg.port)

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.connMaxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, err
}
