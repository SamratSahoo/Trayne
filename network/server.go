package network

import (
	"log"
	"net"
)

const (
	TYPE = "tcp"
)

func InitServer(host string, port string) *net.Listener {
	server, err := net.Listen(TYPE, net.JoinHostPort(host, port))

	if err != nil {
		log.Fatal(err)
	}

	return &server
}
