package main

import (
	"log"

	"github.com/ilbarlo/flavourGeneratorConsumer/pkg/flavourmanager"
)

func main() {
	// Remember: configuration to run into a cluster

	err := flavourmanager.StartConsumer("metrics", "amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("failed to start consumer: %v", err)
	}
}
