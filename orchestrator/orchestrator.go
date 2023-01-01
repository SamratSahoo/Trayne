package orchestrator

import (
	"net"

	network "github.com/SamratSahoo/Trayne/network"
)

var globalConnection *net.Listener = nil

func Start() *net.Listener {
	connection := network.InitServer("localhost", "3000")
	globalConnection = connection
	return connection
}

func Close() {
	(*globalConnection).Close()
}
