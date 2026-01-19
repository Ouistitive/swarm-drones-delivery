package core

import "swarm-drones-delivery/internal/world"

type ClickableEntity interface {
	GetDisplayData() 	string
	Position()			world.Position
}