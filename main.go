package main

import (
	"flag"

	orchestrator "github.com/SamratSahoo/Trayne/orchestrator"
	peripheral "github.com/SamratSahoo/Trayne/peripheral"
	types "github.com/SamratSahoo/Trayne/types"
	utils "github.com/SamratSahoo/Trayne/utils"
)

func main() {
	var nodeType string
	var host string
	var port int
	var peers string
	var peripheralList []string
	flag.StringVar(&nodeType, "type", types.ORCHESTRATOR, "Type of node you want to run (orchestrator or peripheral)")
	flag.StringVar(&peers, "peers", "", "IP addresses of the peripheral nodes you want to connect to")
	flag.StringVar(&host, "host", "localhost", "Host of the node")
	flag.IntVar(&port, "port", -1, "Port to run the node on")
	flag.Parse()

	utils.ParseFlags(&nodeType, &host, &port, &peers, &peripheralList)

	var node types.Node
	if nodeType == types.ORCHESTRATOR {
		node =
			types.Node{
				NodeType:    types.ORCHESTRATOR,
				Server:      orchestrator.InitNode(host, port),
				Host:        host,
				Port:        port,
				Start:       orchestrator.Start,
				Close:       orchestrator.Close,
				Peripherals: &peripheralList,
			}
	} else if nodeType == types.PERIPHERAL {
		node =
			types.Node{
				NodeType: types.PERIPHERAL,
				Server:   peripheral.InitNode(host, port),
				Host:     host,
				Port:     port,
				Start:    peripheral.Start,
				Close:    peripheral.Close,
			}
	}
	defer node.Close(&node)
	node.Start(&node)
}
