package mapgen

import "math/rand"

func weightedRandom(m map[TileState]float64) TileState {
	r := rand.Float64()
	acc := 0.0
	for k, w := range m {
		acc += w
		if r <= acc {
			return k
		}
	}
}
