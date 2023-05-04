package flavourmanager

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var nodeMap = make(NodeMap)

// Function to create RabbitMQ Channel
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

// Function to declare a queue on a RabbitMQ Channel
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

func consumeMsgs(ch *amqp.Channel, qName string) (<-chan amqp.Delivery, error) {
	// si mette in ascolto sulla coda per i messaggi
	msgs, err := ch.Consume(
		qName, // nome della coda
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

func handleMsg(body []byte) {
	var nodeInfo NodeInfo
	err := json.Unmarshal(body, &nodeInfo)
	if err != nil {
		log.Printf("failed to unmarshal JSON: %v", err)
		return
	}

	updateNodeMap(nodeMap, nodeInfo)

}
