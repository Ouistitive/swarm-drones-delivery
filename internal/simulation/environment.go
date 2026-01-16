package simulation

import (
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/objects"
	"swarm-drones-delivery/internal/world"
	"time"
)

type Environment struct {
	agents        []core.IAgent
	spawnedAgents []core.IAgent
	world         *world.Map
	objects       []objects.Delivery

	moveChan   chan core.MoveRequest
	spawnChans []chan core.SpawnRequest
}

func NewEnvironment(m *world.Map) *Environment {
	spawnChans := make([]chan core.SpawnRequest, 0)
	for range len(m.Spawners) {
		spawnChans = append(spawnChans, make(chan core.SpawnRequest))
	}

	return &Environment{
		agents:        make([]core.IAgent, 0),
		spawnedAgents: make([]core.IAgent, 0),
		world:         m,
		objects:       make([]objects.Delivery, 0),
		moveChan:      make(chan core.MoveRequest),
		spawnChans:    spawnChans,
	}
}

func (e *Environment) Start() {
	go e.spawnRequest()
	go e.moveRequest()
	go e.spawnRandomDelivery()
}

func (e *Environment) moveRequest() {
	for moveRequest := range e.moveChan {
		agt := moveRequest.Agt
		agt.Move()
		moveRequest.ResponseChannel <- true
	}
}

func (e *Environment) spawnRequest() {
	for _, spawnChan := range e.spawnChans {
		go func() {
			for spawnRequest := range spawnChan {
				spawnRequest.ResponseChannel <- true
				e.spawnedAgents = append(e.spawnedAgents, spawnRequest.Agt)
				time.Sleep(time.Duration(constants.AGENT_SPAWN_INTERVAL) * time.Millisecond)
			}
		}()
	}
}

func (e *Environment) spawnRandomDelivery() {
	for {
		e.objects = append(e.objects, *objects.NewDelivery("", e.world.RandomPosition()))
		time.Sleep(time.Second)
	}
}

func (e *Environment) AddAgent(factory core.AgentFactory) {
	randomPos, idx := e.world.RandomSpawner()
	e.agents = append(e.agents, factory(randomPos, e.moveChan, e.spawnChans[idx]))
}

func (e *Environment) World() *world.Map {
	return e.world
}

func (e *Environment) Agents() []core.IAgent {
	return e.agents
}

func (e *Environment) SpawnedAgents() []core.IAgent {
	return e.spawnedAgents
}

func (e *Environment) Objects() []objects.Delivery {
	return e.objects
}
