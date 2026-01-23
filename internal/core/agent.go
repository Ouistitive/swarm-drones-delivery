package core

import "swarm-drones-delivery/internal/world"

type AgentID string

type IAgent interface {
	ClickableEntity

	ID() AgentID
	Start()
	Percept()
	Deliberate()
	Act()

	Spawned() bool
	Position() world.Position
	Move()
	Mission() *DeliveryMission
	TargetPos() world.Position

	SurroundingAgents() []IAgent
}

type AgentFactory func(pos world.Position, chanReqs ChannelRequests) IAgent
