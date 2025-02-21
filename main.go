package main

import (
	"context"
	"log"
	"summy/server"
)

func main() {
	s := server.New()
	err := s.StartScheduledTasks(context.Background())
	if err != nil {
		log.Fatal("failed to start scheduled tasks", err)
	}
}
