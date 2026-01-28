package core

import "swarm-drones-delivery/internal/world"

type ChargingMissionRequest struct {
	Agt             IAgent
	ResponseChannel chan ChargingMission
}

type ExitChargingRequest struct {
	Agt             IAgent
	ChargingPoint   *world.ChargingPoint
	ResponseChannel chan bool
}

type MoveRequest struct {
	Agt             IAgent
	ResponseChannel chan bool
}

type SpawnRequest struct {
	Agt             IAgent
	ResponseChannel chan bool
}

type PickRequest struct {
	Agt             IAgent
	Deliv           *Package
	ResponseChannel chan bool
}

type DeliverRequest struct {
	Agt             IAgent
	Deliv           *Package
	ResponseChannel chan bool
}

type PerceptionRequest struct {
	Pos             world.Position
	ResponseChannel chan PerceptionData
}

type ChannelRequests struct {
	ChargingMissionChan  chan ChargingMissionRequest
	ExitChargingChan     chan ExitChargingRequest
	MoveChan             chan MoveRequest
	PickChan             chan PickRequest
	DeliverChan          chan DeliverRequest
	SpawnChan            chan SpawnRequest
	PerceptionChan       chan PerceptionRequest
}

func NewChannelRequests(
	chargingMissionChan chan ChargingMissionRequest,
	exitChargingChan chan ExitChargingRequest,
	moveChan chan MoveRequest,
	pickChan chan PickRequest,
	deliverChan chan DeliverRequest,
	spawnChan chan SpawnRequest,
	perceptionChan chan PerceptionRequest,
) ChannelRequests {
	return ChannelRequests{
		ChargingMissionChan:  chargingMissionChan,
		ExitChargingChan:     exitChargingChan,
		MoveChan:             moveChan,
		PickChan:             pickChan,
		DeliverChan:          deliverChan,
		SpawnChan:            spawnChan,
		PerceptionChan:       perceptionChan,
	}
}
