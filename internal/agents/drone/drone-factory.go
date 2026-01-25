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
		t:               time.Now(),
		env:             env,
		id:              agtId,
		hasSpawned:      false,
		vision:          behaviors.NewVision(constants.VISION_RANGE),
		syncChan:        syncChan,
		requests:        rqs,
		inbox:			 make(chan Info, 10),
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
