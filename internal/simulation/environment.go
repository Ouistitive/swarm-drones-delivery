package simulation

import (
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

type Environment struct {
	nb int
	agents []core.IAgent
	world  *world.Map

	moveChan chan core.MoveRequest
}

func NewEnvironment(m *world.Map) *Environment {
	return &Environment{
		nb: 0,
		agents:   make([]core.IAgent, 0),
		world:    m,
		moveChan: make(chan core.MoveRequest),
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
	e.agents = append(e.agents, factory(e.world.RandomPosition(), e.moveChan))
}

func (e *Environment) World() *world.Map {
	return e.world
}

func (e *Environment) Agents() []core.IAgent {
	return e.agents
}
