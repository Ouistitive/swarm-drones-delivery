package drone

import (
	"math/rand"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"
	"time"
)

func (d *Drone) Deliberate() {
	switch d.state {
	// Generate the next mission to do (recharging or delivering)
	case StateFindingMission:
		d.deliberateFindingMission()
	// Move to a delivery target and grab it if it can
	case StateMovingToDelivery:
		d.deliberateMovingToDelivery()
	// Recharging
	case StateMovingToRecharge:
		d.deliberateMovingToRecharge()
	// Grab the delivery and prepare the next delivery destination
	case StateGrabbing:
		d.deliberateGrabbing()
	// Search random positions based, if the position is found, change the target position
	case StateWandering:
		d.deliberateWandering()
	// Move to delivery position
	case StateMovingToDestination:
		d.deliberateMovingToDestination()
	// Deliver the package at the destination and generate a new mission
	case StateDelivering:
		d.deliberateDelivering()
	}
}

func (d *Drone) deliberateFindingMission() {
	if time.Since(d.t) >= time.Second || d.deliveryMission == nil {
		if d.lastScorePos == (world.Position{}) || utils.GetDistance(d.pos, d.lastScorePos) > constants.RECALC_DISTANCE {
			d.orderBestPackages()
			d.lastScorePos = d.pos
		}

		d.t = time.Now()
		// Try to find a recharge the drone
		if d.battery.Ratio() < constants.BATTERY_EMERGENCY_RATIO && d.state != StateRecharging {
			d.chargingMission = d.getNearestChargingPoint()
			if d.chargingMission != nil {
				d.targetPos = d.chargingMission.TargetCharging.Pos
				d.setDroneStateAndAction(StateMovingToRecharge, ActionMove)
				return
			}
		}

		// If cannot recharge, go deliver a new delivery
		if len(d.allDeliveryMissions) > 0 {
			d.deliveryMission = &d.allDeliveryMissions[0]
		} else {
			d.deliveryMission = nil
		}
		if d.chargingMission == nil && d.deliveryMission != nil && d.deliveryMission.TargetPackage != nil {
			d.targetPos = d.deliveryMission.TargetPackage.Position()
			d.setDroneStateAndAction(StateMovingToDelivery, ActionMove)
		}
	}
}

func (d *Drone) deliberateMovingToDelivery() {
	if d.deliveryMission == nil || !d.deliveryMission.TargetPackage.IsGrabbable() {
		d.setDroneStateAndAction(StateFindingMission, ActionMove)
	} else if d.deliveryMission.TargetPackage != nil && d.isDroneNearTarget(constants.AGENT_CLOSE_DISTANCE) {
		d.setDroneStateAndAction(StateGrabbing, ActionPick)
	}
}

func (d *Drone) deliberateMovingToRecharge() {
	if d.isDroneNearTarget(constants.AGENT_CLOSE_DISTANCE) {
		d.setDroneStateAndAction(StateRecharging, ActionRecharge)
	}
}

func (d *Drone) deliberateGrabbing() {
	if d.deliveryMission.TargetPackage.Carrier == d {
		destPos, exists := d.memory.KnowAddress(&d.deliveryMission.Destination)
		if exists {
			d.targetPos = destPos
			d.setDroneStateAndAction(StateMovingToDestination, ActionMove)
		} else {
			d.targetPos = world.NullPosition()
			d.setDroneStateAndAction(StateWandering, ActionMove)
		}
	} else {
		d.setDroneStateAndAction(StateMovingToDelivery, ActionMove)
	}
}

func (d *Drone) deliberateWandering() {
	if d.targetPos == world.NullPosition() || d.isDroneNearTarget(constants.AGENT_REGENERATION_RANDOM_POS_DISTANCE) {
		d.targetPos = world.NewPosition(rand.Float64()*d.vision.WorldBoundaries.X, rand.Float64()*d.vision.WorldBoundaries.Y)
	}

	p, e := d.memory.KnowAddress(&d.deliveryMission.Destination)
	if e {
		d.targetPos = p
		d.setDroneStateAndAction(StateMovingToDestination, ActionMove)
	}
}

func (d *Drone) deliberateMovingToDestination() {
	if d.isDroneNearTarget(constants.AGENT_CLOSE_DISTANCE) {
		d.setDroneStateAndAction(StateDelivering, ActionDeliver)
	}
}

func (d *Drone) deliberateDelivering() {
	if d.deliveryMission == nil {
		d.setDroneStateAndAction(StateFindingMission, ActionMove)
	} else if d.isDroneNearTarget(constants.AGENT_CLOSE_DISTANCE) {
		d.setDroneStateAndAction(StateDelivering, ActionDeliver)
	}
}
