package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

func main() {

	// Generate a UUID for the node
	nodeId := uuid.New().String()

	// Prompt the user for a human-friendly name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter node name (leave blank for auto): ")
	nameInput, _ := reader.ReadString('\n')
	nodeName := strings.TrimSpace(nameInput)

	if nodeName == "" {
		nodeName = "node-" + nodeId[:8]
	}

	fmt.Println("Nose2Node online")
	fmt.Printf("Node name: %s\n", nodeName)
	fmt.Printf("Node UUID: %s\n", nodeId)
}
