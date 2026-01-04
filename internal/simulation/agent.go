package simulation

import "swarm-drones-delivery/internal/world"

type AgentID string

type IAgent interface {
	Percept()
	Deliberate()
	Act()

	ID() 		AgentID
	Start()
	Position() 	world.Position
	Move()
}

type Agent struct {
	id					AgentID
	syncChan  			chan int

	moveChan 			chan MoveRequest
	moveChanResponse  	chan bool
}

func NewAgent(agtId AgentID, syncChan chan int, moveChan chan MoveRequest) *Agent {
	return &Agent{
		id: agtId,
		syncChan: syncChan,
		moveChan: moveChan,
		moveChanResponse: make(chan bool),
	}
}

func (agt *Agent) ID() AgentID {
	return agt.id
}