package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/SamratSahoo/Trayne/network"
	orchestrator "github.com/SamratSahoo/Trayne/orchestrator"
	peripheral "github.com/SamratSahoo/Trayne/peripheral"
	training "github.com/SamratSahoo/Trayne/training"
	utils "github.com/SamratSahoo/Trayne/utils"
	types "github.com/SamratSahoo/Trayne/utils/types"
)

func main() {
	var nodeType string
	var host string
	var port int
	var peers string
	var peripheralList []string
	var dataDirectory string
	flag.StringVar(&nodeType, "type", types.ORCHESTRATOR, "Type of node you want to run (orchestrator or peripheral)")
	flag.StringVar(&peers, "peers", "", "IP addresses of the peripheral nodes you want to connect to")
	flag.StringVar(&host, "host", "localhost", "Host of the node")
	flag.StringVar(&dataDirectory, "data", "images", "Directory of where your data is located")
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
	} else if nodeType == types.CLIENT {
		message := types.Message{
			MessageType: types.ORCHESTRATOR_TRAINING_INIT,
			Data:        utils.ToMapInterface(training.GetDataset(dataDirectory)),
		}
		network.SendMessage(host, strconv.Itoa(port), message)
	}

	if nodeType != types.CLIENT {
		fmt.Println("============ Node Info ============")
		fmt.Println("Node Type:", node.NodeType)
		fmt.Println("Node Peripherals (if applicable):", node.Peripherals)
		fmt.Println("Node Address:", node.Server.Addr())
		fmt.Println("\n============ Node Logs ============")
		defer node.Close(&node)
		node.Start(&node)
	}
}
