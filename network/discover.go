package network

import (
	"fmt"
	"log"
	"net"

	types "github.com/SamratSahoo/Trayne/types"
)

func FindPeripheralNodes() []string {
	return []string{}
}

func VerifyPeripheralNodes(source string, peerList []string) []string {

	var validPeers []string
	for _, peer := range peerList {
		// Edge case for empty peer list
		if peer == "" {
			continue
		}
		host, port, err := net.SplitHostPort(peer)
		if err != nil {
			log.Fatal(peer, "is not a valid ip address!")
		}
		success, err := SendMessage(host, port, map[string]string{
			"source":      source,
			"destination": net.JoinHostPort(host, port),
			"messageType": types.PERIPHERAL_VERIFICATION,
		})

		if err != nil {
			fmt.Println(err)
		}

		if success {
			validPeers = append(validPeers, peer)
		}

	}
	return validPeers
}
