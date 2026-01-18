package simulation

import (
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"time"
)

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

func (e *Environment) deliverRequest() {
	for deliverRequest := range e.deliverChan {
		del := deliverRequest.Deliv
		if del.State != core.GRABBED {
			deliverRequest.ResponseChannel <- false
		}

		agt := deliverRequest.Agt
		del.State = core.DELIVERED
		del.Carrier = nil
		e.removeMission(*agt.Mission())

		deliverRequest.ResponseChannel <- true
	}
}

func (e *Environment) removeMission(toRemove core.Mission) {
	for i, m := range e.missions {
		if m.Id == toRemove.Id {
			e.missions = append(e.missions[:i], e.missions[i+1:]...)
		}
	}
}

func (e *Environment) generateMissions() {
	for range 3{
		e.missions = append(e.missions, *core.NewMission(core.NewDelivery(e.world.RandomPosition()), e.world.RandomPosition()))
		time.Sleep(time.Second)
	}
}
