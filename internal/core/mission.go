package core

import (
	"swarm-drones-delivery/internal/world"

	"github.com/google/uuid"
)

type MissionType int

const (
	MissionDelivery MissionType = iota
	MissionRecharge
)

type Mission struct {
	Id   uuid.UUID
	Type MissionType

	TargetDelivery *Delivery
	Destination    world.Position
}

func NewDeliveryMission(uuid uuid.UUID, targetDel *Delivery, dest world.Position) *Mission {
	return &Mission{
		Id:             uuid,
		Type:           MissionDelivery,
		TargetDelivery: targetDel,
		Destination:    dest,
	}
}

func NewRechargeMission(uuid uuid.UUID, dest world.Position) *Mission {
	return &Mission{
		Id:             uuid,
		Type:           MissionRecharge,
		Destination:    dest,
		TargetDelivery: nil,
	}
}
