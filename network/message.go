package network

import (
	"encoding/json"
	"log"
	"net"

	types "github.com/SamratSahoo/Trayne/utils/types"
)

func SendMessage(host string, port string, message map[string]interface{}) (bool, error) {
	connection, err := net.Dial(types.CONNECTION_TYPE, net.JoinHostPort(host, port))
	defer connection.Close()
	if err != nil {
		return false, err
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return false, err
	}

	// send message
	_, err = connection.Write(jsonMessage)
	if err != nil {
		return false, err
	}

	return true, nil

}

func DecodeMessage(connection net.Conn) map[string]interface{} {
	// Make a buffer to hold the message
	buffer := TCPReader(1024, &connection)
	// Read the message as a JSON
	var message map[string]interface{}
	err := json.Unmarshal(buffer, &message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}
