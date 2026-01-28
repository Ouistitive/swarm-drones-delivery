package drone

import (
	"sort"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/utils"
)

func (d *Drone) orderBestPackages() {
	scored := make([]core.ScoredMission, len(d.allDeliveryMissions))
	for i, m := range d.allDeliveryMissions {
		// If cannot grab, ignore the package
		if !m.TargetPackage.IsGrabbable() {
			scored[i] = core.ScoredMission{
				Mission: m,
				Score: -1,
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

		scored[i] = core.ScoredMission{
			Mission: m,
			Score: distScore * constants.DISTANCE_PACKAGE_SCORE_WEIGHT +
				destinationDistScore * constants.DISTANCE_DESTINATION_SCORE_WEIGHT,
		}		
	}

	// Rearrange in descending order
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	for i := range scored {
		d.allDeliveryMissions[i] = scored[i].Mission
	}
}