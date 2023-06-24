package flavourmanager

import (
	"fmt"
	"reflect"
)

const (
	Small  Plan = "Small: 11"
	Medium Plan = "Medium: 33"
	Large  Plan = "Large: 66"
)

// Flavour represents a subset of a node's resources
type Flavour struct {
	UID             string     `json:"uid"`
	Name            string     `json:"name"`
	Architecture    string     `json:"architecture"`
	OperatingSystem string     `json:"os"`
	CPUOffer        string     `json:"cpuOffer"`
	MemoryOffer     string     `json:"memoryOffer"`
	PodsOffer       []PodsPlan `json:"podsOffer"`
}

// PodsPlan represents a plan for which is possibile to have a specific amount of available pods
type PodsPlan struct {
	Name      string `json:"name"`
	Available bool   `json:"available"`
	Pods      int64  `json:"availablePods"`
}

// FlavourMap represents a Map of nodes
type FlavourMap map[string]Flavour

// Plan represents a specific Plan for each flavour depending the number of pods
type Plan string

// updateFlavourMap updates the FlavourMap
func updateFlavourMap(flavourMap FlavourMap, newFlavour Flavour) {
	if existingFlavour, ok := flavourMap[newFlavour.UID]; !ok {
		// The node is new, add it to the map
		flavourMap[newFlavour.UID] = newFlavour

		// FOR DEBUG: After the implementation delete this part

		fmt.Printf("\nFlavour: %s \n", newFlavour.Name)
		fmt.Printf("Architecture: %s\n", newFlavour.Architecture)
		fmt.Printf("Operating System: %s\n", newFlavour.OperatingSystem)
		fmt.Printf("CPU: %s\n", newFlavour.CPUOffer)
		fmt.Printf("Memory: %s\n", newFlavour.MemoryOffer)

		fmt.Printf("\nOffer of node %s added!\n", newFlavour.Name)
		upsertFlavour(&newFlavour)
	} else if !reflect.DeepEqual(newFlavour, existingFlavour) {
		// The node already exists, check if there is some changes
		existingFlavour.Name = newFlavour.Name
		existingFlavour.Architecture = newFlavour.Architecture
		existingFlavour.OperatingSystem = newFlavour.OperatingSystem
		existingFlavour.CPUOffer = newFlavour.CPUOffer
		existingFlavour.MemoryOffer = newFlavour.MemoryOffer
		existingFlavour.PodsOffer = newFlavour.PodsOffer
		flavourMap[newFlavour.UID] = existingFlavour

		// FOR DEBUG: After the implementation delete this part

		fmt.Printf("\nFlavour: %s\n", newFlavour.Name)
		fmt.Printf("Architecture: %s\n", newFlavour.Architecture)
		fmt.Printf("Operating System: %s\n", newFlavour.OperatingSystem)
		fmt.Printf("CPU: %s\n", newFlavour.CPUOffer)
		fmt.Printf("Memory: %s\n", newFlavour.MemoryOffer)
		fmt.Printf("PodsOffer:\n")

		fmt.Printf("\nOffer of flavour %s updated!\n", newFlavour.Name)
		upsertFlavour(&newFlavour)
	}
}
