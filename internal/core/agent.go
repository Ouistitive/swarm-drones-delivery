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
	Mission() *Mission
	TargetPos() world.Position

	SurroundingAgents() []IAgent
}

type ChannelRequests struct {
	DeliveryMissionsChan 	chan DeliveryMissionsRequest
	ChargingMissionChan 	chan ChargingMissionRequest
	MoveChan 				chan MoveRequest
	PickChan 				chan PickRequest
	DeliverChan 			chan DeliverRequest
	SpawnChan 				chan SpawnRequest
}

type AgentFactory func(pos world.Position, chanReqs ChannelRequests) IAgent
