package drone

import "swarm-drones-delivery/internal/utils"

func (d *Drone) isDroneNearTarget(dist float64) bool {
	return utils.GetDistance(d.targetPos, d.pos) < dist
}