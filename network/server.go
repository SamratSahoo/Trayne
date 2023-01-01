package network

import (
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "3000"
	TYPE = "tcp"
)

func InitServer() *net.Listener {
	server, err := net.Listen(TYPE, net.JoinHostPort(HOST, PORT))

	if err != nil {
		log.Fatal(err)
	}

	return &server
}
