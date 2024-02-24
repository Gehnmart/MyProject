package config

import "os"

type Backend struct {
	Address string
	Port    int
}

type DB struct {
	DSN string
}

type Config struct {
	Backend Backend
	DB      DB
}

func Init() Config {
	return Config{
		Backend: Backend{
			Address: "127.0.0.1",
			Port:    4000,
		},
		DB: DB{
			DSN: os.Getenv("EASYMESSAGE_DB_DSN"),
		},
	}
}
