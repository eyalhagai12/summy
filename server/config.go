package server

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DbConnectionString string `env:"POSTGRES_URL"`
	GmailClientID      string `env:"GMAIL_CLIENT_ID"`
	GmailClientSecret  string `env:"GMAIL_CLIENT_SECRET"`

	ServerURL string `env:"SERVER_URL"`
}

func LoadServerConfig() Config {
	ctx := context.Background()
	godotenv.Load()

	var config Config
	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal("failed to load server config - ", err)
	}

	return config
}
