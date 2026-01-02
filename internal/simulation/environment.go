package simulation

import (
	"swarm-drones-delivery/internal/agents"
	"swarm-drones-delivery/internal/world"
)

type Environment struct {
	agents 	[]agents.Agent
	Map 	*world.Map
}

func NewEnvironment(m *world.Map) *Environment {
	return &Environment{
		agents: make([]agents.Agent, 0),
		Map: m,
	}
}