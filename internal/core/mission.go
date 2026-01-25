package core

import (
	"fmt"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"

	"github.com/google/uuid"
)

type DeliveryMission struct {
	Id   uuid.UUID

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
		TargetDelivery: targetDel,
		Destination:    dest,
	}
}

func (dm *DeliveryMission) ToString() string {
	if dm.TargetDelivery.Carrier == nil {
		return fmt.Sprintf("Get package to (%s)", utils.PositionToString(dm.TargetDelivery.pos))
	}
	return fmt.Sprintf("Delivery to %s", dm.Destination.Street)
}

func NewRechargeMission(uuid uuid.UUID, dest *world.ChargingPoint, ok bool) *ChargingMission {
	return &ChargingMission{
		Id:             uuid,
		TargetCharging: dest,
		Ok:             ok,
	}
}

func (cm *ChargingMission) ToString() string {
	return fmt.Sprintf("Charging to (%s)", utils.PositionToString(cm.TargetCharging.Pos))
}