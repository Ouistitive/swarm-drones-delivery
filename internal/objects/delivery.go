package objects

import "swarm-drones-delivery/internal/world"

type ObjectId string

type Delivery struct {
	id 	ObjectId
	pos world.Position
}

func NewDelivery(id ObjectId, pos world.Position) *Delivery {
	return &Delivery{
		id: 	id,
		pos: 	pos,
	}
}

func (d *Delivery) Position() world.Position {
	return d.pos
}