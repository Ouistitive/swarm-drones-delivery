package core

import "swarm-drones-delivery/internal/world"

type Mission struct {
    TargetDelivery  *Delivery
    GrabbedDelivery *Delivery
    Destination     world.Position
}

func NewMission(targetDel *Delivery, dest world.Position) *Mission {
	return &Mission{
		TargetDelivery: targetDel,
		Destination: dest,
	}
}