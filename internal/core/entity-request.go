package core

type DeliveryMissionsRequest struct {
	ResponseChannel chan []Mission
}

type ChargingMissionRequest struct {
	Agt 			IAgent
	ResponseChannel chan Mission
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
