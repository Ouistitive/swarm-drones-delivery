package simulation

import (
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
	"time"
)

type Environment struct {
	agents        []core.IAgent
	spawnedAgents []core.IAgent
	world         *world.Map
	objects       []core.Delivery
	missions 	  []core.Mission

	moveChan   	chan core.MoveRequest
	pickchan	chan core.PickRequest
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
		objects:       make([]core.Delivery, 0),
		moveChan:      make(chan core.MoveRequest),
		pickchan: 	   make(chan core.PickRequest),
		spawnChans:    spawnChans,
	}
}

func (e *Environment) Start() {
	go e.spawnRequest()
	go e.moveRequest()
	go e.pickRequest()
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

func (e *Environment) pickRequest() {
	for pickRequest := range e.pickchan {
		del := pickRequest.Deliv
		if del.State == core.GRABBED || del.State == core.DELIVERED {
			pickRequest.ResponseChannel <- false
		}
		agt := pickRequest.Agt
		del.State = core.GRABBED
		del.Carrier = agt
		agt.GrabDelivery(del)
		pickRequest.ResponseChannel <- true
	}
}

func (e *Environment) spawnRandomDelivery() {
	for {
		newDel := core.NewDelivery("", e.world.RandomPosition())
		e.objects = append(e.objects, *newDel)
		e.missions = append(e.missions, *core.NewMission(newDel, e.world.RandomPosition()))
		time.Sleep(time.Second)
	}
}

func (e *Environment) AddAgent(factory core.AgentFactory) {
	randomPos, idx := e.world.RandomSpawner()
	e.agents = append(e.agents, factory(randomPos, e.moveChan, e.pickchan, e.spawnChans[idx]))
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

func (e *Environment) Objects() []core.Delivery {
	return e.objects
}

func (e *Environment) Missions() []core.Mission {
	return e.missions
}