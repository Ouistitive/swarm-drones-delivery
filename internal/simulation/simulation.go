package simulation

import (
	"fmt"
	"swarm-drones-delivery/internal/agents/drone"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
	"sync"
	"time"
)

type Simulation struct {
	Env 		*Environment
	syncChans 	sync.Map
	tickCount 	int
	ticDuration int
}

func NewSimulation(m *world.Map) (*Simulation) {
	sim := &Simulation{
		Env: 			NewEnvironment(m),
		tickCount: 		0,
		ticDuration: 	constants.TIC_DURATION,
	}

	for i := range constants.NB_AGENTS {
		agtId := core.AgentID(fmt.Sprintf("Agent_%d", i))
		syncChan := make(chan int)
		sim.syncChans.Store(agtId, syncChan)
		agtFactory := drone.DroneFactory(sim.Env, agtId, syncChan)
		sim.Env.AddAgent(agtFactory)
	}

	return sim
}

func (s *Simulation) Run() {
	s.Env.Start()

	go func() {
		for {
			s.tickCount++
			time.Sleep(time.Duration(s.ticDuration) * time.Millisecond)
		}
	}()

	for _, agt := range s.Env.Agents() {
		go agt.Start()

		go func(agt core.IAgent) {
			step := 0
			for {
				step++
				syncChan, ok := s.syncChans.Load(agt.ID())
				if !ok {
					fmt.Printf("No sync channel found for agent %s, finishing...\n", agt.ID())
					return
				}

				syncChan.(chan int) <- step
				time.Sleep(time.Millisecond * time.Duration(s.ticDuration))
				<-syncChan.(chan int)
			}
		}(agt)
	}
}