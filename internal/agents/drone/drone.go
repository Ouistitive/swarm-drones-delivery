package drone

import (
	"fmt"
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"
)

type AgentState int

const (
	StateWandering AgentState = iota
	StateMovingToTarget
	StateGrabbing
)

type ActionType int

const (
	ActionMove ActionType = iota
	ActionPick
)

type Drone struct {
	id         core.AgentID
	env        core.IEnvironment
	hasSpawned bool

	vision          behaviors.Vision
	surroundingAgts []core.IAgent

	syncChan  chan int
	moveChan  chan core.MoveRequest
	pickChan  chan core.PickRequest
	spawnChan chan core.SpawnRequest

	pos        world.Position
	targetDir  world.Position
	currentDir world.Position
	velocity   float64

	state      AgentState
	nextAction ActionType

	mission *core.Mission

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
	if d.mission != nil {
		return d.mission.TargetDelivery.Position()
	}
	return world.NewPosition(0, 0)
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
		if time.Since(d.t) >= 5*time.Second || d.mission == nil {
			d.generateTargetPosition()
			d.changeTargetAngle()
			d.t = time.Now()

			if d.mission.TargetDelivery != nil {
				d.setDroneStateAndAction(StateMovingToTarget, ActionMove)
			}
		}
	case StateMovingToTarget:
		d.setDroneStateAndAction(StateMovingToTarget, ActionMove)
		if d.mission.TargetDelivery != nil && utils.GetDistance(d.mission.TargetDelivery.Position(), d.pos) < 0.1 {
			d.setDroneStateAndAction(StateGrabbing, ActionPick)
		}
	case StateGrabbing:
		d.setDroneStateAndAction(StateMovingToTarget, ActionMove)
	}
}

func (d *Drone) Act() {
	switch d.nextAction {
	case ActionMove:
		d.move()
	case ActionPick:
		d.grab()
	}
}