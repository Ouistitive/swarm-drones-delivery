package simulation

import (
	"swarm-drones-delivery/internal/agents"
	Map "swarm-drones-delivery/internal/map"
)

type Environment struct {
	agents 	[]agents.Agent
	Map 	*Map.Map
}