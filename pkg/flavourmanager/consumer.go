package flavourmanager

import (
	"fmt"
	"log"
)

func StartConsumer(queueName string, url string) error {
	// definisce il canale RabbitMQ
	conn, ch, err := createChannel(url)
	if err != nil {
		return fmt.Errorf("failed to create channel: %v", err)
	}
	defer ch.Close()
	defer conn.Close()

	// definisce la coda su cui il consumer ascolterà i messaggi
	err = declareQueue(ch, queueName)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	// si mette in ascolto sulla coda
	msgs, err := consumeMsgs(ch, queueName)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %v", err)
	}

	// loop infinito che legge i messaggi che arrivano sulla coda
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
