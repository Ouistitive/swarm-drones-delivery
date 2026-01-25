package mapgen

import "math/rand"

func isAllTilesCollapsed(m *Map) bool {
	for i := range m.Height {
		for j := range m.Width {
			if len(m.Tiles[i][j].PossibleStates) > 1 {
				return false
			}
		}
	}

	return true
}

func getPosWithLowestEntropy(m *Map) ([][]int) {
	positions := make([][]int, 0)
	minEntropy := 100

	for i := range m.Height {
		for j := range m.Width {
			if minEntropy > len(m.Tiles[i][j].PossibleStates) && len(m.Tiles[i][j].PossibleStates) > 1 {
				minEntropy = len(m.Tiles[i][j].PossibleStates)
			}
		}
	}

	for i := range m.Height {
		for j := range m.Width {
			if minEntropy == len(m.Tiles[i][j].PossibleStates) {
				positions = append(positions, []int{i, j})
			}
		}
	}

	return positions
}

func collapse(m *Map, x, y int) {
	state := m.Tiles[x][y].PossibleStates[rand.Intn(len(m.Tiles[x][y].PossibleStates))]
	m.Tiles[x][y].PossibleStates = []TileState{state}
	m.Tiles[x][y].FinalState = &state
}

func WaveFunctionCollapse(width, height int) Map {
	m := NewMap(width, height)
	m.Tiles[0][5].PossibleStates = []TileState{TileChargingPoint, TileDestination}

	for !isAllTilesCollapsed(&m) {
		lowestEntropy := getPosWithLowestEntropy(&m)
		var collapsedPosition []int
		if len(lowestEntropy) == 0 {
			collapsedPosition = lowestEntropy[0]
		} else {
			collapsedPosition = lowestEntropy[rand.Intn(len(lowestEntropy))]	
		}

		collapse(&m, collapsedPosition[0], collapsedPosition[1])
	}

	return m
}
