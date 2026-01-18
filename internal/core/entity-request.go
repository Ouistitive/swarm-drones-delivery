package core

type MoveRequest struct {
	Agt IAgent
	ResponseChannel chan bool	
}

type SpawnRequest struct {
	Agt IAgent
	ResponseChannel chan bool
}

type PickRequest struct {
	Agt IAgent
	Deliv *Delivery
	ResponseChannel chan bool
}

type DeliverRequest struct {
	Agt IAgent
	Deliv *Delivery
	ResponseChannel chan bool
}