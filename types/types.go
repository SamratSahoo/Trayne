package utils

import "net"

const (
	ORCHESTRATOR = "orchestrator"
	PERIPHERAL   = "peripheral"
)

const (
	PERIPHERAL_VERIFICATION = "peripheral verification"
	PERIPHERAL_CONFIRMATION = "peripheral confirmation"
)

const (
	CONNECTION_TYPE = "tcp"
)

type Node struct {
	NodeType    string
	Server      net.Listener
	Host        string
	Port        int
	Start       func(*Node)
	Close       func(*Node)
	Peripherals *[]string
}
