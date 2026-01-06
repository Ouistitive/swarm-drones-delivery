package utils

import (
	"math"
	"swarm-drones-delivery/internal/world"
)

func GetDistance(pos1, pos2 world.Position) float64 {
	return math.Sqrt(math.Pow(pos1.X - pos2.X, 2) + math.Pow(pos1.Y - pos2.Y, 2))
}