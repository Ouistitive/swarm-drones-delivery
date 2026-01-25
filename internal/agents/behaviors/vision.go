package behaviors

import (
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

func (v *Vision) IsAgentDetected(currAgt *world.Position, pos *world.Position) bool {
	return utils.GetDistance(*currAgt, *pos) <= v.visionRange
}
