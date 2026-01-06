package core

import "swarm-drones-delivery/internal/world"

type AgentID string

type IAgent interface {
	Percept()
	Deliberate()
	Act()

	ID() 		AgentID
	Start()
	Position() 	world.Position
	Move()

	SurroundingAgents() []IAgent
}

type AgentFactory func(pos world.Position, moveChan chan MoveRequest) IAgent