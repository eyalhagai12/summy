package server

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DbConnectionString string `env:"POSTGRES_URL,required"`

	GmailClientID     string `env:"GMAIL_CLIENT_ID"`
	GmailClientSecret string `env:"GMAIL_CLIENT_SECRET"`

	HostURL string `env:"SERVER_URL"`

	WorkerPoolSize      int `env:"WORKER_POOL_SIZE,default=10"`
	TaskBufferPerWorker int `env:"WORKER_BUFFER,default=20"`
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
