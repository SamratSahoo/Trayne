package types

import "net"

const (
	ORCHESTRATOR = "orchestrator"
	PERIPHERAL   = "peripheral"
	CLIENT       = "client"
)

const (
	PERIPHERAL_VERIFICATION    = "peripheral verification"
	PERIPHERAL_CONFIRMATION    = "peripheral confirmation"
	ORCHESTRATOR_TRAINING_INIT = "orchestrator training init"
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
