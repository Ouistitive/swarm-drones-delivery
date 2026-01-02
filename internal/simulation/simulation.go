package simulation

import "swarm-drones-delivery/internal/world"

type Simulation struct {
	Env *Environment
}

func NewSimulation(m *world.Map) (*Simulation) {
	return &Simulation{
		Env: NewEnvironment(m),
	}
}