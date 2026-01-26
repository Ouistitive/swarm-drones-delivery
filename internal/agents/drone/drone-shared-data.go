package drone

import (
	"swarm-drones-delivery/internal/world"
)

type DroneSharedData struct {
	Addresses []world.DeliveryDestination
}
