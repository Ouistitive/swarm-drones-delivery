package drone

import (
	"sort"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/utils"
	"swarm-drones-delivery/internal/world"
)

type DroneSharedData struct {
	Addresses []world.DeliveryDestination
}

type ScoredMission struct {
	mission core.DeliveryMission
	score   float64
}

func (d *Drone) orderBestPackages() {
	scored := make([]ScoredMission, len(d.allDeliveryMissions))
	for i, m := range d.allDeliveryMissions {
		// If cannot grab, ignore the package
		if !m.TargetPackage.IsGrabbable() {
			scored[i] = ScoredMission{
				mission: m,
				score: -1,
			}
			continue
		}

		// Based on the distance
		distScore := utils.NormMin(
			utils.GetDistance(d.pos, m.TargetPackage.Position()),
			0,
			d.vision.WorldBoundaries.X,
		)

		// Based on the distance between the package and destination (if position is knowed by the drone)
		destinationDistScore := 0.0
		if p, e := d.memory.KnowAddress(&m.Destination); e {
			destinationDistScore = utils.NormMin(
				utils.GetDistance(p, m.TargetPackage.Position()),
				0,
				d.vision.WorldBoundaries.X,
			)
		}

		scored[i] = ScoredMission{
			mission: m,
			score: distScore * 0.2 +
				destinationDistScore * 0.1,
		}		
	}

	// Rearrange in descending order
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score > scored[j].score
	})

	for i := range scored {
		d.allDeliveryMissions[i] = scored[i].mission
	}
}