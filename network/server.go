package network

import (
	"log"
	"net"

	types "github.com/SamratSahoo/Trayne/types"
)

func InitServer(host string, port string) net.Listener {
	server, err := net.Listen(types.CONNECTION_TYPE, net.JoinHostPort(host, port))

	if err != nil {
		log.Fatal(err)
	}

	return server
}

func EndServer(node *types.Node) {
	node.Server.Close()
}
