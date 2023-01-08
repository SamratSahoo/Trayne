package orchestrator

import (
	"fmt"

	"github.com/SamratSahoo/Trayne/training"
	types "github.com/SamratSahoo/Trayne/utils/types"
)

func messageRouter(message map[string]interface{}) {
	switch messageType := message["messageType"]; messageType {
	case types.ORCHESTRATOR_TRAINING_INIT:
		processedData := training.TypeAssertDataset(message["data"])
		fmt.Println(training.PartitionDataset(processedData, 2))
	}
}
