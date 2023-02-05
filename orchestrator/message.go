package orchestrator

import (
	"encoding/json"
	"net"

	"github.com/SamratSahoo/Trayne/network"
	training "github.com/SamratSahoo/Trayne/training"
	"github.com/SamratSahoo/Trayne/utils"
	types "github.com/SamratSahoo/Trayne/utils/types"
)

func messageRouter(message types.Message, node *types.Node) {
	switch messageType := message.MessageType; messageType {
	case types.ORCHESTRATOR_TRAINING_INIT:
		dataset := map[string][]string{}
		// Process JSON Types Properly
		bytes, _ := json.Marshal(message.Data)
		json.Unmarshal(bytes, &dataset)
		splits := training.PartitionDataset(dataset, len(*node.Peripherals))

		// Send split data to peripheral nodes to start training
		for index, peripheral := range *node.Peripherals {
			host, port, _ := net.SplitHostPort(peripheral)
			message := types.Message{
				MessageType: types.PERIPHERAL_TRAINING,
				Data:        utils.ToMapInterface(splits[index]),
			}
			network.SendMessage(host, port, message)
		}
	}
}
