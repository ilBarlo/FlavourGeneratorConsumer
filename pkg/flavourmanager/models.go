package flavourmanager

import (
	"fmt"
	"reflect"
)

// NodeInfo represents a node and its resources
type NodeInfo struct {
	UID             string          `json:"uid"`
	Name            string          `json:"name"`
	Architecture    string          `json:"architecture"`
	OperatingSystem string          `json:"os"`
	ResourceMetrics ResourceMetrics `json:"resources"`
}

// ResourceMetrics represents resources of a certain node
type ResourceMetrics struct {
	CPUTotal        string `json:"totalCPU"`
	CPUAvailable    string `json:"availableCPU"`
	MemoryTotal     string `json:"totalMemory"`
	MemoryAvailable string `json:"availableMemory"`
}

// NodeMap represents a Map of nodes
type NodeMap map[string]NodeInfo

// updateNodeMap updates the NodeMap
func updateNodeMap(nodeMap NodeMap, newNode NodeInfo) {
	if existingNode, ok := nodeMap[newNode.UID]; !ok {
		// The node is new, add it to the map
		nodeMap[newNode.UID] = newNode
		fmt.Printf("Metrics of node %s added!\n", newNode.Name)
	} else if !reflect.DeepEqual(newNode, existingNode) {
		// The node already exists, check if there is some changes
		existingNode.Name = newNode.Name
		existingNode.Architecture = newNode.Architecture
		existingNode.OperatingSystem = newNode.OperatingSystem
		existingNode.ResourceMetrics = newNode.ResourceMetrics
		nodeMap[newNode.UID] = existingNode
		fmt.Printf("Metrics of node %s updated!\n", newNode.Name)
	}
}
