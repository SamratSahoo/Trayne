package peripheral

import (
	"net"

	network "github.com/SamratSahoo/Trayne/network"
	types "github.com/SamratSahoo/Trayne/types"
)

func InitNode() net.Listener {
	connection := network.InitServer("localhost", "4000")
	// Any other code to initialize peripheral node here
	return connection
}
func Start() {}

func Close(node *types.Node) {
	network.EndServer(node)
	// Peripheral Node Cleanup Code Here
}
