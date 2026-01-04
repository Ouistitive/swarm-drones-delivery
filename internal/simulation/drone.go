package simulation

import (
	"math/rand"
	"swarm-drones-delivery/internal/world"
)

type Drone struct {
	*Agent

	pos world.Position
}

func (d *Drone) Start() {
	var step int
	for {
		step = <-d.syncChan
		d.Percept()
		d.Deliberate()
		d.Act()
		step++
		d.syncChan <- step
	}
}

func (d *Drone) Percept() {
	return
}

func (d *Drone) Deliberate() {
	return
}

func (d *Drone) Act() {
	d.moveChan <- MoveRequest{Agt: d, ResponseChannel: d.moveChanResponse}
	<-d.moveChanResponse
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

func NewDrone(agtId AgentID, pos world.Position, syncChan chan int, moveChan chan MoveRequest) *Drone {
	return &Drone{
		Agent: NewAgent(agtId, syncChan, moveChan),
		pos:   pos,
	}
}

func (d *Drone) Position() world.Position {
	return d.pos
}
