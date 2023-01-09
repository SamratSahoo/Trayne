package orchestrator

import (
	"encoding/json"
	"net"

	"github.com/SamratSahoo/Trayne/network"
	training "github.com/SamratSahoo/Trayne/training"
	types "github.com/SamratSahoo/Trayne/utils/types"
)

func messageRouter(message map[string]interface{}, node *types.Node) {
	switch messageType := message["messageType"]; messageType {
	case types.ORCHESTRATOR_TRAINING_INIT:
		dataset := map[string][]string{}
		// Process JSON Types Properly
		bytes, _ := json.Marshal(message["data"])
		json.Unmarshal(bytes, &dataset)
		splits := training.PartitionDataset(dataset, len(*node.Peripherals))

		// Send split data to peripheral nodes to start training
		for index, peripheral := range *node.Peripherals {
			host, port, _ := net.SplitHostPort(peripheral)
			network.SendMessage(host, port, map[string]interface{}{
				"messageType": types.PERIPHERAL_TRAINING,
				"data":        splits[index],
			})
		}
	}
}
