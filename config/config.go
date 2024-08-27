package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server
	Postgres
}

type Server struct {
	Host     string
	Port     string
	LogLevel string
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func New() *Config {
	if os.Getenv("ENV") != "docker" {
		if err := godotenv.Load(); err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	}

	return &Config{
		Server: Server{
			Host:     os.Getenv("SERVER_HOST"),
			Port:     os.Getenv("SERVER_PORT"),
			LogLevel: os.Getenv("SERVER_LOG_LEVEL"),
		},
		Postgres: Postgres{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
	}
}
