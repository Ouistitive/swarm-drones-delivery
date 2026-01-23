package drone

import (
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

func NewDrone(env core.IEnvironment, agtId core.AgentID, pos world.Position, syncChan chan int, deliveryMissionsChan chan core.DeliveryMissionsRequest, chargingMissionChan chan core.ChargingMissionRequest, moveChan chan core.MoveRequest, pickChan chan core.PickRequest, deliverChan chan core.DeliverRequest, spawnChan chan core.SpawnRequest) *Drone {
	return &Drone{
		t:                    time.Now(),
		env:                  env,
		id:                   agtId,
		hasSpawned:           false,
		vision:               behaviors.NewVision(constants.VISION_RANGE),
		syncChan:             syncChan,
		deliveryMissionsChan: deliveryMissionsChan,
		chargingMissionsChan: chargingMissionChan,
		moveChan:             moveChan,
		pickChan:             pickChan,
		deliverChan:          deliverChan,
		spawnChan:            spawnChan,
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
			chanReqs.DeliveryMissionsChan,
			chanReqs.ChargingMissionChan,
			chanReqs.MoveChan,
			chanReqs.PickChan,
			chanReqs.DeliverChan,
			chanReqs.SpawnChan,
		)
	}
}
