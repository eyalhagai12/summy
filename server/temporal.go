package server

import (
	"log"

	"go.temporal.io/sdk/client"
)

func ConnectToTemporal() client.Client {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	return c
}
