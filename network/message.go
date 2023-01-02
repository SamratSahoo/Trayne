package network

import (
	"encoding/json"
	"log"
	"net"

	types "github.com/SamratSahoo/Trayne/types"
)

func SendMessage(host string, port string, message map[string]string) {
	connection, err := net.Dial(types.CONNECTION_TYPE, net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal(err)
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	// send message
	_, err = connection.Write(jsonMessage)
	if err != nil {
		log.Fatal(err)
	}

}
