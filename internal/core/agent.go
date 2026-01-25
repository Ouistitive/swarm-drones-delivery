package core

import "swarm-drones-delivery/internal/world"

type AgentID string

type AgentView struct {
	ID  AgentID
	Pos world.Position
}

func NewAgentView(id AgentID, pos world.Position) AgentView {
	return AgentView{
		ID: 	id,
		Pos: 	pos,
	}
}

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

	SurroundingAgents() []AgentView
}

type AgentFactory func(pos world.Position, chanReqs ChannelRequests) IAgent
