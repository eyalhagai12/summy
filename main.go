package main

import (
	"log"
	"summy/server"
)

func main() {
	config := server.LoadServerConfig()
	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal("failed to start scheduled tasks", err)
	}
}
