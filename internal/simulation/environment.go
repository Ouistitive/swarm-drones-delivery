package simulation

import (
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

type Environment struct {
	agents        []core.IAgent
	spawnedAgents []core.IAgent
	world         *world.Map
	objects       []core.Delivery
	missions      []core.DeliveryMission
	destinations  []world.DeliveryDestination

	deliveryMissionsChan chan core.DeliveryMissionsRequest
	chargingMissionChan  chan core.ChargingMissionRequest
	exitChargingChan     chan core.ExitChargingRequest
	moveChan             chan core.MoveRequest
	pickchan             chan core.PickRequest
	deliverChan          chan core.DeliverRequest
	perceptionChan       chan core.PerceptionRequest
	spawnChans           []chan core.SpawnRequest
}

func NewEnvironment(w *world.Map) *Environment {
	spawnChans := make([]chan core.SpawnRequest, 0)
	for range len(w.Spawners) {
		spawnChans = append(spawnChans, make(chan core.SpawnRequest))
	}

	var dests []world.DeliveryDestination
	for _, d := range w.DeliveryDests {
		dests = append(dests, d)
	}

	return &Environment{
		agents:               make([]core.IAgent, 0),
		spawnedAgents:        make([]core.IAgent, 0),
		world:                w,
		destinations: 		  dests,
		objects:              make([]core.Delivery, 0),
		deliveryMissionsChan: make(chan core.DeliveryMissionsRequest),
		chargingMissionChan:  make(chan core.ChargingMissionRequest),
		exitChargingChan:     make(chan core.ExitChargingRequest),
		moveChan:             make(chan core.MoveRequest),
		pickchan:             make(chan core.PickRequest, 50),
		deliverChan:          make(chan core.DeliverRequest, 50),
		spawnChans:           spawnChans,
		perceptionChan:       make(chan core.PerceptionRequest, 0),
	}
}

func (e *Environment) Start() {
	go e.spawnRequest()
	go e.moveRequest()
	go e.pickRequest()
	go e.perceptionRequest()
	go e.deliverRequest()
	go e.deliveryMissionsRequest()
	go e.chargingMissionRequest()
	go e.exitChargingRequest()
	go e.generateMissions()
}

func (e *Environment) AddAgent(factory core.AgentFactory) {
	randomPos, idx := e.world.RandomSpawner()
	chanRqs := core.NewChannelRequests(e.deliveryMissionsChan, e.chargingMissionChan, e.exitChargingChan, e.moveChan, e.pickchan, e.deliverChan, e.spawnChans[idx], e.perceptionChan)
	e.agents = append(e.agents, factory(randomPos, chanRqs))
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

func (e *Environment) Destinations() []world.DeliveryDestination {
	return e.destinations
}

func (e *Environment) Missions() []core.DeliveryMission {
	return e.missions
}
