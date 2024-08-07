package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}

// Need to declare this variable through terminal
// ex export DATABASE_URL=postgres://user:password@localhost:5432/mydb
