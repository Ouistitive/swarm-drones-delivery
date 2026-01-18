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
	Mission() 	*Mission
	GrabDelivery(del *Delivery)
	TargetPos() world.Position

	SurroundingAgents() []IAgent
}

type AgentFactory func(pos world.Position, moveChan chan MoveRequest, pickChan chan PickRequest, deliverChan chan DeliverRequest, spawnChan chan SpawnRequest) IAgent