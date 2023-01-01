package network

import (
	"log"
	"net"

	types "github.com/SamratSahoo/Trayne/types"
)

const (
	TYPE = "tcp"
)

func InitServer(host string, port string) net.Listener {
	server, err := net.Listen(TYPE, net.JoinHostPort(host, port))

	if err != nil {
		log.Fatal(err)
	}

	return server
}

func EndServer(node *types.Node) {
	node.Server.Close()
}
