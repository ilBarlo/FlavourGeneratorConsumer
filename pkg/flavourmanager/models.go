package flavourmanager

import (
	"fmt"
	"reflect"
)

// Info of the node
type NodeInfo struct {
	UID             string          `json:"uid"`
	Name            string          `json:"name"`
	Architecture    string          `json:"architecture"`
	OperatingSystem string          `json:"os"`
	ResourceMetrics ResourceMetrics `json:"resources"`
}

// Metrics of the resource of a certain node
type ResourceMetrics struct {
	CPUTotal        string `json:"totalCPU"`
	CPUAvailable    string `json:"availableCPU"`
	MemoryTotal     string `json:"totalMemory"`
	MemoryAvailable string `json:"availableMemory"`
}

// Map of the nodes
type NodeMap map[string]NodeInfo

// Update the NodeMap
func updateNodeMap(nodeMap NodeMap, newNode NodeInfo) {
	if existingNode, ok := nodeMap[newNode.UID]; ok {
		// The node already exists, check if there is some changes
		if ok && reflect.DeepEqual(newNode, existingNode) {
			// Nothing
		} else {
			existingNode.Name = newNode.Name
			existingNode.Architecture = newNode.Architecture
			existingNode.OperatingSystem = newNode.OperatingSystem
			existingNode.ResourceMetrics = newNode.ResourceMetrics
			nodeMap[newNode.UID] = existingNode
			fmt.Printf("Metrics of node %s updated!\n", newNode.Name)
		}
	} else {
		// The node is new, add it to the map
		nodeMap[newNode.UID] = newNode
		fmt.Printf("Metrics of node %s added!\n", newNode.Name)
	}
}
