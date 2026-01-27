package core

import (
	"swarm-drones-delivery/internal/world"
)

type PerceptionData struct {
	Agents       []AgentView
	Destinations []world.DeliveryDestination
	Missions 	 []DeliveryMission
}
