package core

import (
	"swarm-drones-delivery/internal/world"

	"github.com/google/uuid"
)

type Mission struct {
	Id 				uuid.UUID
    TargetDelivery  *Delivery
    Destination     world.Position
}

func NewMission(uuid uuid.UUID, targetDel *Delivery, dest world.Position) *Mission {
	return &Mission{
		Id: 			uuid,
		TargetDelivery: targetDel,
		Destination: 	dest,
	}
}