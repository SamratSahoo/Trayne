package peripheral

import (
	"fmt"

	types "github.com/SamratSahoo/Trayne/utils/types"
)

func messageRouter(message types.Message) {
	switch messageType := message.MessageType; messageType {
	case types.PERIPHERAL_VERIFICATION:
		// do not need to actually do anything here
		// used more as an endpoint for the orchestrator to verify the connection
		// kept here as a filler case for clarification
		fmt.Println("NODE: node has been verified by", message.Data["source"])
	case types.PERIPHERAL_TRAINING:
		fmt.Println("NODE: node has received data from orchestrator")
	}
}
