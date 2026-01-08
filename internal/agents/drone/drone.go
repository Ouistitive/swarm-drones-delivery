package drone

import (
	"fmt"
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

type Drone struct {
	id  core.AgentID
	env core.IEnvironment

	vision          behaviors.Vision
	surroundingAgts []core.IAgent

	syncChan         chan int
	moveChan         chan core.MoveRequest
	moveChanResponse chan bool

	pos 			world.Position
	targetPos		world.Position
	targetDir 		world.Position
	currentDir 		world.Position
	velocity 		float64

	t time.Time

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

func (d *Drone) Start() {
	fmt.Println("Drone started:", d.id)

	for {
		step := <-d.syncChan
		d.Percept()
		d.Deliberate()
		d.Act()
		d.syncChan <- step + 1
	}
}

func (d *Drone) Percept() {
	agts := d.env.Agents()
	d.surroundingAgts = d.surroundingAgts[:0]

	for _, a := range agts {
		if d.vision.IsAgentDetected(d.pos, a) {
			d.surroundingAgts = append(d.surroundingAgts, a)
		}
	}
}

func (d *Drone) Deliberate() {
	if time.Since(d.t) >= 3 * time.Second {
		d.generateTargetPosition()
		d.changeTargetAngle()
		
		d.t = time.Now()
	}
}

func (d *Drone) Act() {
	moveChanResponse := make(chan bool)
	d.moveChan <- core.MoveRequest{Agt: d, ResponseChannel: moveChanResponse}
	<- moveChanResponse
}