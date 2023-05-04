package flavourmanager

import (
	"fmt"
	"log"
)

// Function that start a new Consumer on a RabbitMQ Channel
func StartConsumer(queueName string, url string) error {

	conn, ch, err := createChannel(url)
	if err != nil {
		return fmt.Errorf("failed to create channel: %v", err)
	}
	defer ch.Close()
	defer conn.Close()

	err = declareQueue(ch, queueName)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	msgs, err := consumeMsgs(ch, queueName)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %v", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			handleMsg(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
