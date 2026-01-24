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

type DeliveryMission struct {
	Id   uuid.UUID
	Type MissionType

	TargetDelivery *Delivery
	Destination    world.Address
	Ok             bool
}

type ChargingMission struct {
	Id uuid.UUID

	TargetCharging *world.ChargingPoint
	Ok             bool
}

func NewDeliveryMission(uuid uuid.UUID, targetDel *Delivery, dest world.Address) *DeliveryMission {
	return &DeliveryMission{
		Id:             uuid,
		Type:           MissionDelivery,
		TargetDelivery: targetDel,
		Destination:    dest,
	}
}

func NewRechargeMission(uuid uuid.UUID, dest *world.ChargingPoint, ok bool) *ChargingMission {
	return &ChargingMission{
		Id:             uuid,
		TargetCharging: dest,
		Ok:             ok,
	}
}
