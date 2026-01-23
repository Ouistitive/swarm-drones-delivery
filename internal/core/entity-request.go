package core

import "swarm-drones-delivery/internal/world"

type DeliveryMissionsRequest struct {
	ResponseChannel chan []DeliveryMission
}

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
	Deliv           *Delivery
	ResponseChannel chan bool
}

type DeliverRequest struct {
	Agt             IAgent
	Deliv           *Delivery
	ResponseChannel chan bool
}

type ChannelRequests struct {
	DeliveryMissionsChan chan DeliveryMissionsRequest
	ChargingMissionChan  chan ChargingMissionRequest
	ExitChargingChan     chan ExitChargingRequest
	MoveChan             chan MoveRequest
	PickChan             chan PickRequest
	DeliverChan          chan DeliverRequest
	SpawnChan            chan SpawnRequest
}

func NewChannelRequests(
	deliveryMissionsChan chan DeliveryMissionsRequest,
	chargingMissionChan chan ChargingMissionRequest,
	exitChargingChan chan ExitChargingRequest,
	moveChan chan MoveRequest,
	pickChan chan PickRequest,
	deliverChan chan DeliverRequest,
	spawnChan chan SpawnRequest,
) ChannelRequests {
	return ChannelRequests{
		DeliveryMissionsChan: deliveryMissionsChan,
		ChargingMissionChan:  chargingMissionChan,
		ExitChargingChan: 	  exitChargingChan,
		MoveChan:             moveChan,
		PickChan:             pickChan,
		DeliverChan:          deliverChan,
		SpawnChan:            spawnChan,
	}
}
