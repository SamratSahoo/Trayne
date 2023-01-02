package orchestrator

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	network "github.com/SamratSahoo/Trayne/network"
	types "github.com/SamratSahoo/Trayne/types"
)

func InitNode(host string, port int) net.Listener {
	server := network.InitServer(host, strconv.Itoa(port))
	// Any other code to initialize orchestrator node here
	return server
}

func Start(node *types.Node) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		node.Close(node)
		os.Exit(1)
	}()

	for {
		// Listen for a connection
		connection, err := node.Server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		connectionHandler(connection)

	}
}

func connectionHandler(connection net.Conn) {

}

func Close(node *types.Node) {
	// Orchestrator Node Cleanup Code Here
	network.EndServer(node)
}
