package drone

import (
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

func NewDrone(env core.IEnvironment, agtId core.AgentID, pos world.Position, syncChan chan int, moveChan chan core.MoveRequest, pickChan chan core.PickRequest, spawnChan chan core.SpawnRequest) *Drone {
	return &Drone{
		t:               time.Now(),
		env:             env,
		id:              agtId,
		hasSpawned:      false,
		vision:          behaviors.NewVision(constants.VISION_RANGE),
		syncChan:        syncChan,
		moveChan:        moveChan,
		pickChan: 		 pickChan,
		spawnChan:       spawnChan,
		pos:             pos,
		surroundingAgts: []core.IAgent{},
		targetDir:       world.NullPosition(),
		currentDir:      world.NullPosition(),
		velocity:        0.0,
		state:           StateWandering,
	}
}

func DroneFactory(
	env core.IEnvironment,
	agtId core.AgentID,
	syncChan chan int,
) core.AgentFactory {
	return func(pos world.Position, moveChan chan core.MoveRequest, pickChan chan core.PickRequest, spawnChan chan core.SpawnRequest) core.IAgent {
		return NewDrone(
			env,
			agtId,
			pos,
			syncChan,
			moveChan,
			pickChan,
			spawnChan,
		)
	}
}
