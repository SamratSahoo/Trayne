package main

import (
	"flag"
	"fmt"
	"strings"

	network "github.com/SamratSahoo/Trayne/network"
	orchestrator "github.com/SamratSahoo/Trayne/orchestrator"
	peripheral "github.com/SamratSahoo/Trayne/peripheral"
	types "github.com/SamratSahoo/Trayne/types"
)

func main() {
	var nodeType string
	var peripheralNodes string
	var peripheralList []string
	flag.StringVar(&nodeType, "type", types.ORCHESTRATOR, "Type of node you want to run (orchestrator or peripheral)")
	flag.StringVar(&peripheralNodes, "peers", "", "IP addresses of the peripheral nodes you want to connect to")
	flag.Parse()

	if nodeType == types.ORCHESTRATOR {
		peripheralList := strings.Split(peripheralNodes, " ")
		if peripheralList == nil {
			peripheralList = network.FindPeripheralNodes()
		}
		peripheralList = network.VerifyPeripheralNodes(peripheralList)
	}

	var node types.Node

	if nodeType == types.ORCHESTRATOR {
		node =
			types.Node{
				NodeType:    types.ORCHESTRATOR,
				Server:      orchestrator.InitNode(),
				Start:       orchestrator.Start,
				Close:       orchestrator.Close,
				Peripherals: &peripheralList,
			}
	} else if nodeType == types.PERIPHERAL {
		node =
			types.Node{
				NodeType: types.PERIPHERAL,
				Server:   peripheral.InitNode(),
				Start:    peripheral.Start,
				Close:    peripheral.Close,
			}
	}
	defer node.Close(&node)

	fmt.Println(node.Server.Addr())
}
