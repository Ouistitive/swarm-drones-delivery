package drone

import (
	"fmt"
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

//go:generate stringer -type=AgentState
type AgentState int

const (
	StateFindingMission AgentState = iota
	StateWandering
	StateMovingToDelivery
	StateMovingToRecharge
	StateMovingToDestination

	StateGrabbing
	StateDelivering
	StateRecharging
)

//go:generate stringer -type=ActionType
type ActionType int

const (
	ActionMove ActionType = iota
	ActionPick
	ActionDeliver
	ActionRecharge
)

type Drone struct {
	id         core.AgentID
	env        core.IEnvironment
	hasSpawned bool

	vision          behaviors.Vision
	surroundingAgts []core.AgentView

	syncChan chan int
	requests core.ChannelRequests

	pos        world.Position
	targetPos  world.Position // What the drone is trying to go in a current state
	targetDir  world.Position
	currentDir world.Position
	velocity   float64

	state      AgentState
	nextAction ActionType

	deliveryMission *core.DeliveryMission
	chargingMission *core.ChargingMission
	battery         behaviors.Battery
	memory          behaviors.Memory

	t time.Time
}

func (d *Drone) Spawned() bool {
	return d.hasSpawned
}

func (d *Drone) SurroundingAgents() []core.AgentView {
	return d.surroundingAgts
}

func (d *Drone) ID() core.AgentID {
	return d.id
}

func (d *Drone) Position() world.Position {
	return d.pos
}

func (d *Drone) TargetPos() world.Position {
	return d.targetPos
}

func (d *Drone) GetDisplayData() string {
	text := fmt.Sprintf("AgentID: %s\nState: %s\nAction: %s\nBattery: %d", d.id, d.state.String(), d.nextAction.String(), int(d.battery.Ratio()*100))
	return text
}

func (d *Drone) Mission() *core.DeliveryMission {
	return d.deliveryMission
}

func (d *Drone) Start() {
	fmt.Println("Drone started:", d.id)

	for !d.hasSpawned {
		startChanResponse := make(chan bool)
		d.requests.SpawnChan <- core.SpawnRequest{Agt: d, ResponseChannel: startChanResponse}
		d.hasSpawned = <-startChanResponse
	}

	for {
		step := <-d.syncChan
		d.Percept()
		d.Deliberate()
		d.Act()
		d.syncChan <- step + 1
	}
}

func (d *Drone) Percept() {
	data := d.getPerceptionData()
	d.surroundingAgts = d.surroundingAgts[:0]
	for _, a := range data.Agents {
		if d.vision.IsAgentDetected(&d.pos, &a.Pos) {
			d.surroundingAgts = append(d.surroundingAgts, a)
		}
	}

	for _, de := range data.Destinations {
		if d.vision.IsAgentDetected(&d.pos, &de.Pos) {
			d.memory.AddAddress(&de.Address, &de.Pos)
		}
	}
}

func (d *Drone) Deliberate() {
	switch d.state {
	// Generate the next mission to do (recharging or delivering)
	case StateFindingMission:
		if time.Since(d.t) >= time.Second || d.deliveryMission == nil {
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
			d.deliveryMission = d.getMissions()
			if d.chargingMission == nil && d.deliveryMission != nil && d.deliveryMission.TargetDelivery != nil {
				d.targetPos = d.deliveryMission.TargetDelivery.Position()
				d.setDroneStateAndAction(StateMovingToDelivery, ActionMove)
			}
		}
	// Move to a delivery target and grab it if it can
	case StateMovingToDelivery:
		if d.deliveryMission == nil || !d.deliveryMission.TargetDelivery.IsGrabbable() {
			d.setDroneStateAndAction(StateFindingMission, ActionMove)
		} else if d.deliveryMission.TargetDelivery != nil && d.isDroneNearTarget(constants.AGENT_CLOSE_DISTANCE) {
			d.setDroneStateAndAction(StateGrabbing, ActionPick)
		}
	// Recharging
	case StateMovingToRecharge:
		if d.isDroneNearTarget(constants.AGENT_CLOSE_DISTANCE) {
			d.setDroneStateAndAction(StateRecharging, ActionRecharge)
		}
	// Grab the delivery and prepare the next delivery destination
	case StateGrabbing:
		if d.deliveryMission.TargetDelivery.Carrier == d {
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
	// Search random positions based, if the position is found, change the target position
	case StateWandering:
		if d.targetPos == world.NullPosition() || d.isDroneNearTarget(constants.AGENT_REGENERATION_RANDOM_POS_DISTANCE) {
			d.targetPos = d.env.World().RandomPosition()
		}

		p, e := d.memory.KnowAddress(&d.deliveryMission.Destination)
		if e {
			d.targetPos = p
			d.setDroneStateAndAction(StateMovingToDestination, ActionMove)
		}
	// Move to delivery position
	case StateMovingToDestination:
		if d.isDroneNearTarget(constants.AGENT_CLOSE_DISTANCE) {
			d.setDroneStateAndAction(StateDelivering, ActionDeliver)
		}
	// Deliver the package at the destination and generate a new mission
	case StateDelivering:
		if d.deliveryMission == nil {
			d.setDroneStateAndAction(StateFindingMission, ActionMove)
		} else if d.isDroneNearTarget(constants.AGENT_CLOSE_DISTANCE) {
			d.setDroneStateAndAction(StateDelivering, ActionDeliver)
		}
	}
}

func (d *Drone) Act() {
	switch d.nextAction {
	case ActionMove:
		if d.battery.Consume(constants.BATTERY_DISCHARGING_MOVE) {
			d.move()
		}
	case ActionPick:
		if d.battery.Consume(constants.BATTERY_DISCHARGING_PICK) {
			d.grab()
		}
	case ActionDeliver:
		if d.battery.Consume(constants.BATTERY_DISCHARGING_DELIVER) {
			d.deliver()
		}
	case ActionRecharge:
		d.battery.Recharge(constants.BATTERY_CHARGING_RATE)
		if d.battery.IsFull() {
			d.deliveryMission = nil
			d.exitCharging()
			d.setDroneStateAndAction(StateFindingMission, ActionMove)
		}
	}
}
