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
)

type Drone struct {
	id         core.AgentID
	env        core.IEnvironment
	hasSpawned bool

	vision          behaviors.Vision
	surroundingAgts []core.IAgent

	syncChan  		chan int
	missionsChan 	chan core.MissionsRequest
	moveChan  		chan core.MoveRequest
	pickChan  		chan core.PickRequest
	deliverChan 	chan core.DeliverRequest
	spawnChan 		chan core.SpawnRequest

	pos        world.Position
	targetPos  world.Position // What the drone is trying to go in a current state
	targetDir  world.Position
	currentDir world.Position
	velocity   float64

	state      AgentState
	nextAction ActionType

	mission *core.Mission
	battery behaviors.Battery

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
	text:= fmt.Sprintf("AgentID: %s\nState: %s\nAction: %s\nBattery: %d", d.id, d.state.String(), d.nextAction.String(), int(d.battery.Ratio() * 100))
	return text
}

func (d *Drone) Mission() *core.Mission {
	return d.mission
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
		if time.Since(d.t) >= time.Second || d.mission == nil {
			d.generateTargetPosition()
			d.changeTargetAngle()
			d.t = time.Now()

			if d.mission != nil && d.mission.TargetDelivery != nil {
				d.targetPos = d.mission.TargetDelivery.Position()
				d.setDroneStateAndAction(StateMovingToDelivery, ActionMove)
			}
		}
	case StateMovingToDelivery:
		if d.mission == nil || !d.mission.TargetDelivery.IsGrabbable() {
			d.setDroneStateAndAction(StateWandering, ActionMove)
		} else if d.mission.TargetDelivery != nil && utils.GetDistance(d.mission.TargetDelivery.Position(), d.pos) < 0.1 {
			d.setDroneStateAndAction(StateGrabbing, ActionPick)
		}
	case StateGrabbing:
		if d.mission.TargetDelivery.Carrier == d {
			d.targetPos = d.mission.Destination
			d.setDroneStateAndAction(StateMovingToDestination, ActionMove)	
		} else {
			d.setDroneStateAndAction(StateMovingToDelivery, ActionMove)
		}
	case StateMovingToDestination:
		if utils.GetDistance(d.mission.Destination, d.pos) < 0.1 {
			d.setDroneStateAndAction(StateDelivering, ActionDeliver)
		}
	case StateDelivering:
		if d.mission == nil {
			d.setDroneStateAndAction(StateWandering, ActionMove)
		} else if utils.GetDistance(d.mission.Destination, d.pos) < 0.1 {
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
	}
}