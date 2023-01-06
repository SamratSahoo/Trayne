package orchestrator

import (
	"fmt"

	types "github.com/SamratSahoo/Trayne/utils/types"
)

func messageRouter(message map[string]interface{}) {
	switch messageType := message["messageType"]; messageType {
	case types.ORCHESTRATOR_TRAINING_INIT:
		fmt.Println("CLIENT: a user has initialized the training")
	}
}
