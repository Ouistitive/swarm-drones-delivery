package simulation

import (
	"fmt"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

type Environment struct {
	agents        []core.IAgent
	spawnedAgents []core.IAgent
	world         *world.Map
	objects       []core.Delivery
	missions 	  []core.Mission

	moveChan   	chan core.MoveRequest
	pickchan	chan core.PickRequest
	deliverChan	chan core.DeliverRequest
	spawnChans []chan core.SpawnRequest
}

func NewEnvironment(w *world.Map) *Environment {
	spawnChans := make([]chan core.SpawnRequest, 0)
	for range len(w.Spawners) {
		spawnChans = append(spawnChans, make(chan core.SpawnRequest))
	}

	return &Environment{
		agents:        make([]core.IAgent, 0),
		spawnedAgents: make([]core.IAgent, 0),
		world:         w,
		objects:       make([]core.Delivery, 0),
		moveChan:      make(chan core.MoveRequest),
		pickchan: 	   make(chan core.PickRequest),
		deliverChan:   make(chan core.DeliverRequest),
		spawnChans:    spawnChans,
	}
}

func (e *Environment) Start() {
	go e.spawnRequest()
	go e.moveRequest()
	go e.pickRequest()
	go e.deliverRequest()
	go e.generateMissions()
}

func (e *Environment) AddAgent(factory core.AgentFactory) {
	randomPos, idx := e.world.RandomSpawner()
	e.agents = append(e.agents, factory(randomPos, e.moveChan, e.pickchan, e.deliverChan, e.spawnChans[idx]))
}

func (e *Environment) World() *world.Map {
	fmt.Println(e.missions)
	return e.world
}

func (e *Environment) Agents() []core.IAgent {
	return e.agents
}

func (e *Environment) SpawnedAgents() []core.IAgent {
	return e.spawnedAgents
}

func (e *Environment) Missions() []core.Mission {
	return e.missions
}