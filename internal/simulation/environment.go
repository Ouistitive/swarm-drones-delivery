package simulation

import (
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

type Environment struct {
	Agents 			[]core.IAgent
	Map    			*world.Map

	moveChan		chan core.MoveRequest
}

func NewEnvironment(m *world.Map) *Environment {
	return &Environment{
		Agents: 	make([]core.IAgent, 0),
		Map:    	m,
		moveChan:	make(chan core.MoveRequest),
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

func (e *Environment) AddAgent(factory core.AgentFactory) {
	e.Agents = append(e.Agents, factory(e.Map.RandomPosition(), e.moveChan))
}
