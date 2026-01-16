package core

import (
	"swarm-drones-delivery/internal/objects"
	"swarm-drones-delivery/internal/world"
)

type IEnvironment interface {
	Agents() 		[]IAgent
	SpawnedAgents()	[]IAgent
	World() 		*world.Map
	Objects()		[]objects.Delivery
}