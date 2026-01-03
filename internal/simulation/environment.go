package simulation

import (
	"math/rand"
	"swarm-drones-delivery/internal/world"
)

type MoveRequest struct {
	Agt IAgent
	ResponseChannel chan bool
}

type Environment struct {
	Agents 			[]IAgent
	Map    			*world.Map

	moveChan		chan MoveRequest
}

func NewEnvironment(m *world.Map) *Environment {
	return &Environment{
		Agents: 	make([]IAgent, 0),
		Map:    	m,
		moveChan:	make(chan MoveRequest),
	}
}

func (e *Environment) Start() {
	go e.moveRequest()
}

func (e *Environment) moveRequest() {
	for moveRequest := range e.moveChan {
		agt := moveRequest.Agt
		agt.Move()
		moveRequest.ResponseChannel <- true
	}
}

func (e *Environment) AddAgent(droneId AgentID, syncChan chan int) {
	e.Agents = append(e.Agents, NewDrone(
		droneId, 
		world.NewPosition(
			rand.Float64() * float64(e.Map.Width-1), 
			rand.Float64() * float64(e.Map.Height-1),
		),
		syncChan,
		e.moveChan,
	))
}
