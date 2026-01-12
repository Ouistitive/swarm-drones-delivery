package core

import "swarm-drones-delivery/internal/world"

type AgentID string

type IAgent interface {
	ID() 		AgentID
	Start()
	Percept()
	Deliberate()
	Act()

	Spawned() 	bool
	Position() 	world.Position
	Move()
	TargetPos() world.Position

	SurroundingAgents() []IAgent
}

type AgentFactory func(pos world.Position, moveChan chan MoveRequest, spawnChan chan SpawnRequest) IAgent