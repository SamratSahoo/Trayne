package utils

import "net"

const (
	ORCHESTRATOR = "orchestrator"
	PERIPHERAL   = "peripheral"
)

type Node struct {
	NodeType    string
	Server      net.Listener
	Start       func()
	Close       func(*Node)
	Peripherals *[]string
}
