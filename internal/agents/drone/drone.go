package drone

import (
	"fmt"
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"
)

//go:generate stringer -type=AgentState
type AgentState int

const (
	StateWandering AgentState = iota
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
	surroundingAgts []core.IAgent

	syncChan             chan int
	deliveryMissionsChan chan core.DeliveryMissionsRequest
	chargingMissionsChan chan core.ChargingMissionRequest
	moveChan             chan core.MoveRequest
	pickChan             chan core.PickRequest
	deliverChan          chan core.DeliverRequest
	spawnChan            chan core.SpawnRequest
	exitChargingChan     chan core.ExitChargingRequest

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

	t time.Time
}

func (d *Drone) Spawned() bool {
	return d.hasSpawned
}

func (d *Drone) SurroundingAgents() []core.IAgent {
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
		d.spawnChan <- core.SpawnRequest{Agt: d, ResponseChannel: startChanResponse}
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
	agts := d.env.SpawnedAgents()
	d.surroundingAgts = d.surroundingAgts[:0]

	for _, a := range agts {
		if d.vision.IsAgentDetected(d.pos, a) {
			d.surroundingAgts = append(d.surroundingAgts, a)
		}
	}
}

func (d *Drone) Deliberate() {
	switch d.state {
	case StateWandering:
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
			d.generateTargetPosition()
			if d.chargingMission == nil && d.deliveryMission != nil && d.deliveryMission.TargetDelivery != nil {
				d.targetPos = d.deliveryMission.TargetDelivery.Position()
				d.setDroneStateAndAction(StateMovingToDelivery, ActionMove)
			}
		}
	case StateMovingToDelivery:
		if d.deliveryMission == nil || !d.deliveryMission.TargetDelivery.IsGrabbable() {
			d.setDroneStateAndAction(StateWandering, ActionMove)
		} else if d.deliveryMission.TargetDelivery != nil && utils.GetDistance(d.deliveryMission.TargetDelivery.Position(), d.pos) < 0.1 {
			d.setDroneStateAndAction(StateGrabbing, ActionPick)
		}
	case StateMovingToRecharge:
		if utils.GetDistance(d.targetPos, d.pos) < 0.1 {
			d.setDroneStateAndAction(StateRecharging, ActionRecharge)
		}
	case StateGrabbing:
		if d.deliveryMission.TargetDelivery.Carrier == d {
			d.targetPos = d.deliveryMission.Destination
			d.setDroneStateAndAction(StateMovingToDestination, ActionMove)
		} else {
			d.setDroneStateAndAction(StateMovingToDelivery, ActionMove)
		}
	case StateMovingToDestination:
		if utils.GetDistance(d.deliveryMission.Destination, d.pos) < 0.1 {
			d.setDroneStateAndAction(StateDelivering, ActionDeliver)
		}
	case StateDelivering:
		if d.deliveryMission == nil {
			d.setDroneStateAndAction(StateWandering, ActionMove)
		} else if utils.GetDistance(d.deliveryMission.Destination, d.pos) < 0.1 {
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
			d.setDroneStateAndAction(StateWandering, ActionMove)
		}
	}
}
