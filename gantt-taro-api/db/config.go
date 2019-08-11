package db

import (
	"os"
)

type Config struct {
	Host     string
	User     string
	Password string
}

func NewDBConfig() *Config {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")

	return &Config{
		Host:     host,
		User:     user,
		Password: pass,
	}
}
