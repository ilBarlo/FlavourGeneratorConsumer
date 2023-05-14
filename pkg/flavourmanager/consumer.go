package flavourmanager

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

// StartConsumer starts a new Consumer on a RabbitMQ Channel
func StartConsumer(subject1 string, url string) error {

	nc, err := connectNATS(url)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to the topic that manage flavours
	sub, err := nc.Subscribe(subject1, func(msg *nats.Msg) {
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
