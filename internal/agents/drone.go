package agents

import (
	"fmt"
	"math/rand"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

type Drone struct {
	id					core.AgentID
	syncChan  			chan int

	moveChan 			chan core.MoveRequest
	moveChanResponse  	chan bool

	pos world.Position
}

func DroneFactory(
	agtId core.AgentID,
	syncChan chan int,
) core.AgentFactory {
	return func(pos world.Position, moveChan chan core.MoveRequest) core.IAgent {
		return NewDrone(
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
	
}

func (d *Drone) Deliberate() {
	
}

func (d *Drone) Act() {
	moveChanResponse := make(chan bool)
	d.moveChan <- core.MoveRequest{Agt: d, ResponseChannel: moveChanResponse}
	<-moveChanResponse
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

func NewDrone(agtId core.AgentID, pos world.Position, syncChan chan int, moveChan chan core.MoveRequest) *Drone {
	return &Drone{
		id: agtId, 
		syncChan: syncChan,
		moveChan: moveChan,
		pos:   pos,
		moveChanResponse: make(chan bool),
	}
}

func (d *Drone) Position() world.Position {
	return d.pos
}
