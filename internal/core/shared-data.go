package core

import (
	"swarm-drones-delivery/internal/world"
)

type AgentMessage struct {
	Addresses []world.DeliveryDestination
}
