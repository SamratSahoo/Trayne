package orchestrator

import (
	"net"

	network "github.com/SamratSahoo/Trayne/network"
	types "github.com/SamratSahoo/Trayne/types"
)

func InitNode() net.Listener {
	connection := network.InitServer("localhost", "3000")
	// Any other code to initialize orchestrator node here
	return connection
}

func Start() {}

func Close(node *types.Node) {
	network.EndServer(node)
	// Orchestrator Node Cleanup Code Here
}
