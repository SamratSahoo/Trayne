package main

import (
	"fmt"
	"net"
	"os"

	orchestrator "github.com/SamratSahoo/Trayne/orchestrator"
	peripheral "github.com/SamratSahoo/Trayne/peripheral"
)

const (
	ORCHESTRATOR = "orchestrator"
	PERIPHERAL   = "peripheral"
)

func main() {
	nodeType := os.Args[1]
	var globalNode *net.Listener = nil

	if nodeType == ORCHESTRATOR {
		globalNode = orchestrator.Start()
		defer orchestrator.Close()
	} else if nodeType == PERIPHERAL {
		globalNode = peripheral.Start()
		defer peripheral.Close()
	}

	fmt.Println((*globalNode).Addr())
}
