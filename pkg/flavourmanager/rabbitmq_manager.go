package flavourmanager

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var flavourMap = make(FlavourMap)

// createChannel creates a RabbitMQ Channel
func createChannel(url string) (*amqp.Connection, *amqp.Channel, error) {
	// Connection to the server RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	// Create Channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	return conn, ch, nil
}

// declareQueue declares a queue on a RabbitMQ Channel
func declareQueue(ch *amqp.Channel, queueName string) error {

	_, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

// consumeMsgs consumes messages arrived on a RabbitMQ Channel
func consumeMsgs(ch *amqp.Channel, qName string) (<-chan amqp.Delivery, error) {

	msgs, err := ch.Consume(
		qName, // queue name
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}

	return msgs, nil
}

// handleMsg handles the message received
func handleMsg(body []byte) {
	var flavour Flavour
	err := json.Unmarshal(body, &flavour)
	if err != nil {
		log.Printf("failed to unmarshal JSON: %v", err)
		return
	}

	updateFlavourMap(flavourMap, flavour)

}
