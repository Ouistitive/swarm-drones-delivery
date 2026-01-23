package simulation

import (
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"
	"time"

	"github.com/google/uuid"
)

func (e *Environment) deliveryMissionsRequest() {
	for missionRequest := range e.deliveryMissionsChan {
		cp := make([]core.Mission, len(e.missions))
		copy(cp, e.missions)
		missionRequest.ResponseChannel <- cp
	}
}

func (e *Environment) chargingMissionRequest() {
	for missionRequest := range e.chargingMissionChan {
		agtPos := missionRequest.Agt.Position()
		nearestChargingPoint := e.getNearestChargingPoint(agtPos)
		if nearestChargingPoint != nil {
			nearestChargingPoint.State = world.ChargingPointReserved
			missionRequest.ResponseChannel <- *core.NewRechargeMission(uuid.New(), nearestChargingPoint.Pos, true)
		} else {
			missionRequest.ResponseChannel <- *core.NewRechargeMission(uuid.New(), world.NewPosition(-1, -1), false)
		}
	}
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
		del.SetPosition(agt.Position())
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
	for {
		e.missions = append(e.missions, *core.NewDeliveryMission(uuid.New(), core.NewDelivery(e.world.RandomWarehouses()), e.world.RandomDeliveryDestination()))
		time.Sleep(time.Second)
	}
}

func (e *Environment) getNearestChargingPoint(agtPos world.Position) *world.ChargingPoint {
	nearestDist := 1000.0
	var nearestChargingPoint *world.ChargingPoint = nil
	
	for i := range e.world.ChargingPoints {
		cp := &e.world.ChargingPoints[i]

		if cp.State == world.ChargingPointFree {
			dist := utils.GetDistance(agtPos, cp.Pos)
			if dist < nearestDist {
				nearestDist = dist
				nearestChargingPoint = cp
			}
		}
	}

	return nearestChargingPoint
}