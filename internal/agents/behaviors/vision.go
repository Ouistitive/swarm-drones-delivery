package behaviors

import (
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"
)

type Vision struct {
	visionRange float64
}

func NewVision(visionRange float64) Vision {
	return Vision{
		visionRange: visionRange,
	}
}

func (v *Vision) IsAgentDetected(currAgt world.Position, agt core.IAgent) bool {
	return utils.GetDistance(currAgt, agt.Position()) <= v.visionRange
}
