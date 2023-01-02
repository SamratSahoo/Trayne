package network

import (
	"log"
	"net"

	types "github.com/SamratSahoo/Trayne/types"
)

func FindPeripheralNodes() []string {
	return []string{}
}

func VerifyPeripheralNodes(source string, peerList []string) []string {
	for _, peer := range peerList {
		// Edge case for empty peer list
		if peer == "" {
			continue
		}
		host, port, err := net.SplitHostPort(peer)
		if err != nil {
			log.Fatal(peer, "is not a valid ip address!")
		}
		SendMessage(host, port, map[string]string{
			"source":      source,
			"messageType": types.PERIPHERAL_VERIFICATION,
		})
	}
	return peerList
}
