package main

import (
	"log"

	"github.com/ilbarlo/flavourGeneratorConsumer/pkg/flavourmanager"
	"github.com/nats-io/nats.go"
)

const natsURL = nats.DefaultURL

func main() {
	// Remember: configuration to run into a cluster

	err := flavourmanager.StartConsumer("flavours", natsURL)
	if err != nil {
		log.Fatalf("failed to start consumer: %v", err)
	}
}
