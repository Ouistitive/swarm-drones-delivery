package core

import (
	"fmt"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"

	"github.com/google/uuid"
)

type DeliveryMission struct {
	Id uuid.UUID

	TargetPackage *Package
	Destination   world.Address
	Ok            bool
}

type ScoredMission struct {
	Mission DeliveryMission
	Score   float64
}

type ChargingMission struct {
	Id uuid.UUID

	TargetCharging *world.ChargingPoint
	Ok             bool
}

func NewDeliveryMission(uuid uuid.UUID, targetDel *Package, dest world.Address) *DeliveryMission {
	return &DeliveryMission{
		Id:            uuid,
		TargetPackage: targetDel,
		Destination:   dest,
	}
}

func (dm *DeliveryMission) ToString() string {
	if dm.TargetPackage.Carrier == nil {
		return fmt.Sprintf("Get package to (%s)", utils.PositionToString(dm.TargetPackage.pos))
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
