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
