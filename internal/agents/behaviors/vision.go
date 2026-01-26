package behaviors

import (
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"
)

type Vision struct {
	WorldBoundaries world.Position
	visionRange 	float64
}

func NewVision(visionRange float64, w, h float64) Vision {
	return Vision{
		visionRange: 		visionRange,
		WorldBoundaries: 	world.NewPosition(w, h),
	}
}

func (v *Vision) IsAgentDetected(currAgt *world.Position, pos *world.Position) bool {
	return utils.GetDistance(*currAgt, *pos) <= v.visionRange
}
