package drone

import (
	"time"

	"swarm-drones-delivery/internal/agents/behaviors"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"
)

func NewDrone(env core.IEnvironment, agtId core.AgentID, pos world.Position, syncChan chan int, moveChan chan core.MoveRequest) *Drone {
	return &Drone{
		t: time.Now(),
		env:              env,
		id:               agtId,
		vision:           behaviors.NewVision(constants.VISION_RANGE),
		syncChan:         syncChan,
		moveChan:         moveChan,
		pos:              pos,
		moveChanResponse: make(chan bool),
		surroundingAgts:  []core.IAgent{},
		targetDir: 		  world.NullPosition(),
		currentDir: 	  world.NullPosition(),
		targetPos: 		  env.World().RandomPosition(),
		velocity: 		  0.0,
	}
}

func DroneFactory(
	env core.IEnvironment,
	agtId core.AgentID,
	syncChan chan int,
) core.AgentFactory {
	return func(pos world.Position, moveChan chan core.MoveRequest) core.IAgent {
		return NewDrone(
			env,
			agtId,
			pos,
			syncChan,
			moveChan,
		)
	}
}