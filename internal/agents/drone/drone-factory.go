package drone

import (
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

func NewDrone(env core.IEnvironment, agtId core.AgentID, pos world.Position, syncChan chan int, rqs core.ChannelRequests) *Drone {
	return &Drone{
		t:                    time.Now(),
		env:                  env,
		id:                   agtId,
		hasSpawned:           false,
		vision:               behaviors.NewVision(constants.VISION_RANGE),
		syncChan:             syncChan,
		deliveryMissionsChan: rqs.DeliveryMissionsChan,
		chargingMissionsChan: rqs.ChargingMissionChan,
		exitChargingChan: 	  rqs.ExitChargingChan,
		moveChan:             rqs.MoveChan,
		pickChan:             rqs.PickChan,
		deliverChan:          rqs.DeliverChan,
		spawnChan:            rqs.SpawnChan,
		pos:                  pos,
		surroundingAgts:      []core.IAgent{},
		targetDir:            world.NullPosition(),
		currentDir:           world.NullPosition(),
		velocity:             0.0,
		state:                StateWandering,
		battery:              behaviors.NewBattery(),
	}
}

func DroneFactory(
	env core.IEnvironment,
	agtId core.AgentID,
	syncChan chan int,
) core.AgentFactory {
	return func(pos world.Position, chanReqs core.ChannelRequests) core.IAgent {
		return NewDrone(
			env,
			agtId,
			pos,
			syncChan,
			chanReqs,
		)
	}
}
