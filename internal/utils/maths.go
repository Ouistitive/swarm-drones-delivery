package utils

import (
	"fmt"
	"math"
	"strconv"
	"swarm-drones-delivery/internal/world"
)

func GetDistance(pos1, pos2 world.Position) float64 {
	return math.Sqrt(math.Pow(pos1.X-pos2.X, 2) + math.Pow(pos1.Y-pos2.Y, 2))
}

func FloatToString(f float64, dec int) string {
	return strconv.FormatFloat(f, 'f', dec, 64)
}

func PositionToString(pos world.Position) string {
	return fmt.Sprintf("%s, %s", FloatToString(pos.X, 0), FloatToString(pos.Y, 0))
}