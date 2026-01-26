package drone

import (
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

func NewDrone(agtId core.AgentID, worldBoundaries, pos world.Position, syncChan chan int, rqs core.ChannelRequests) *Drone {
	return &Drone{
		t:               time.Now(),
		id:              agtId,
		hasSpawned:      false,
		vision:          behaviors.NewVision(constants.VISION_RANGE, worldBoundaries.X, worldBoundaries.Y),
		syncChan:        syncChan,
		requests:        rqs,
		inbox:           make(chan DroneSharedData, 10),
		pos:             pos,
		surroundingAgts: []core.AgentView{},
		targetDir:       world.NullPosition(),
		currentDir:      world.NullPosition(),
		velocity:        0.0,
		state:           StateFindingMission,
		battery:         behaviors.NewBattery(),
		memory:          behaviors.NewMemory(),
	}
}

func DroneFactory(
	agtId core.AgentID,
	syncChan chan int,
	worldBoundaries world.Position,
) core.AgentFactory {
	return func(pos world.Position, chanReqs core.ChannelRequests) core.IAgent {
		return NewDrone(
			agtId,
			worldBoundaries,
			pos,
			syncChan,
			chanReqs,
		)
	}
}
