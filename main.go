package main

import (
	"flag"
	"fmt"
	"net"
	"strings"

	network "github.com/SamratSahoo/Trayne/network"
	orchestrator "github.com/SamratSahoo/Trayne/orchestrator"
	peripheral "github.com/SamratSahoo/Trayne/peripheral"
)

const (
	ORCHESTRATOR = "orchestrator"
	PERIPHERAL   = "peripheral"
)

func main() {
	var nodeType string
	var peripheralNodes string

	flag.StringVar(&nodeType, "type", ORCHESTRATOR, "Type of node you want to run (orchestrator or peripheral)")
	flag.StringVar(&peripheralNodes, "peers", "", "IP addresses of the peripheral nodes you want to connect to")
	flag.Parse()

	peripheralList := strings.Split(peripheralNodes, " ")
	if peripheralList == nil {
		peripheralList = network.FindPeripheralNodes()
	}
	network.VerifyPeripheralNodes(peripheralList)

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
