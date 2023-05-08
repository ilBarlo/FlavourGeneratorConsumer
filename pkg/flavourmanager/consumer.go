package flavourmanager

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

// Function that start a new Consumer on a RabbitMQ Channel
func StartConsumer(subject string, url string) error {

	nc, err := connectNATS(url)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to a topic
	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		handleMsg(msg.Data)
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	log.Println("Consumer started. Waiting for messages...")

	// Wait for SIGINT or SIGTERM to gracefully shutdown the application
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down consumer...")

	return nil
}
