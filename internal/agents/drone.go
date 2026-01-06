package agents

import (
	"fmt"
	"math/rand"
	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
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

	pos world.Position
}

func (d *Drone) SurroundingAgents() []core.IAgent {
	return d.surroundingAgts
}

func DroneFactory(
	env core.IEnvironment,
	agtId core.AgentID,
	syncChan chan int,
) core.AgentFactory {
	return func(pos world.Position, moveChan chan core.MoveRequest) core.IAgent {
		return NewDrone(
			env,
			agtId,
			pos,
			syncChan,
			moveChan,
		)
	}
}

func (d *Drone) ID() core.AgentID {
	return d.id
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

}

func (d *Drone) Act() {
	moveChanResponse := make(chan bool)
	d.moveChan <- core.MoveRequest{Agt: d, ResponseChannel: moveChanResponse}
	<- moveChanResponse
}

func (d *Drone) Move() {
	dir := rand.Intn(4)

	switch dir {
	case 0:
		d.pos.X += 0.1
	case 1:
		d.pos.X -= 0.1
	case 2:
		d.pos.Y += 0.1
	case 3:
		d.pos.Y -= 0.1
	}
}

func NewDrone(env core.IEnvironment, agtId core.AgentID, pos world.Position, syncChan chan int, moveChan chan core.MoveRequest) *Drone {
	return &Drone{
		env:              env,
		id:               agtId,
		vision:           behaviors.NewVision(constants.VISION_RANGE),
		syncChan:         syncChan,
		moveChan:         moveChan,
		pos:              pos,
		moveChanResponse: make(chan bool),
		surroundingAgts:  []core.IAgent{},
	}
}

func (d *Drone) Position() world.Position {
	return d.pos
}
