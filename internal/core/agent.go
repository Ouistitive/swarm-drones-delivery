package core

import "swarm-drones-delivery/internal/world"

type AgentID string

type IAgent interface {
	ClickableEntity
	
	ID() 		AgentID
	Start()
	Percept()
	Deliberate()
	Act()

	Spawned() 	bool
	Position() 	world.Position
	Move()
	Mission() 	*Mission
	TargetPos() world.Position

	SurroundingAgents() []IAgent
}

type AgentFactory func(pos world.Position, missionsChan chan MissionsRequest, moveChan chan MoveRequest, pickChan chan PickRequest, deliverChan chan DeliverRequest, spawnChan chan SpawnRequest) IAgent