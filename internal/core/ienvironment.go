package core

import "swarm-drones-delivery/internal/world"

type IEnvironment interface {
	Agents() 		[]IAgent
	SpawnedAgents()	[]IAgent
	World() 		*world.Map
}