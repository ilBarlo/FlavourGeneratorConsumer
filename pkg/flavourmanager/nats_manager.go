package flavourmanager

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

// flavourMap is a map of Flavours
var flavourMap = make(FlavourMap)

// connectNATS creates a connection to a NATS server
func connectNATS(url string) (*nats.Conn, error) {
	// Connect to NATS server
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return nc, nil
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
