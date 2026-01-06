package core

import "swarm-drones-delivery/internal/world"

type IEnvironment interface {
	Agents() 	[]IAgent
	World() 	*world.Map
}