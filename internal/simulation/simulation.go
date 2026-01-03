package simulation

import (
	"fmt"
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
		ticDuration: 	50,
	}

	for i := range 20 {
		agtId := AgentID(fmt.Sprintf("Agent_%d", i))

		syncChan := make(chan int)
		sim.syncChans.Store(agtId, syncChan)
		sim.Env.AddAgent(agtId, syncChan)
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

	for _, agt := range s.Env.Agents {
		go agt.Start()

		go func(agt IAgent) {
			step := 0
			for {
				step++
				c, ok := s.syncChans.Load(agt.ID())
				if !ok {
					fmt.Printf("No sync channel found for agent %s, finishing...\n", agt.ID())
					return
				}

				c.(chan int) <- step
				time.Sleep(1 * time.Millisecond * time.Duration(s.ticDuration))
				<-c.(chan int)
			}
		}(agt)
	}
}