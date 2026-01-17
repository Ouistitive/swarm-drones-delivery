package core

import (
	"swarm-drones-delivery/internal/world"
)

type ObjectId string

type DeliveryState int

const (
	FREE 	DeliveryState = iota
	GRABBED
	DELIVERED
)

type Delivery struct {
	id 		ObjectId
	pos 	world.Position
	State 	DeliveryState
	Carrier IAgent
}

func NewDelivery(id ObjectId, pos world.Position) *Delivery {
	return &Delivery{
		id: 	id,
		pos: 	pos,
		State: 	FREE,
	}
}

func (d *Delivery) Position() world.Position {
	if d.Carrier != nil && d.State == GRABBED {
		return d.Carrier.Position()
	}
	return d.pos
}