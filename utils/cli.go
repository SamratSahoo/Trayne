package utils

import (
	"net"
	"strconv"
	"strings"

	"github.com/SamratSahoo/Trayne/network"
	types "github.com/SamratSahoo/Trayne/types"
)

func ParseFlags(nodeType *string, host *string, port *int, peers *string, peripheralList *[]string) {
	// Process the Orchestrator CLI Arguments
	if *nodeType == types.ORCHESTRATOR {
		if *port == -1 {
			*port = 3000
		}

		*peripheralList = strings.Split(*peers, " ")
		if peripheralList == nil {
			*peripheralList = network.FindPeripheralNodes()
		}
		*peripheralList = network.VerifyPeripheralNodes(
			net.JoinHostPort(*host, strconv.Itoa(*port)),
			*peripheralList,
		)

	}

	// Process the Peripheral CLI Arguments
	if *nodeType == types.PERIPHERAL {
		if *port == -1 {
			*port = 4000
		}

	}
}
