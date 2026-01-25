package mapgen

import (
	"os"
	"strings"
)

func WriteMap(m *Map) {
	text := ToString(m)
	os.WriteFile("data/maps/generated_layout.txt", []byte(text), 0644)
}

func ToString(m *Map) string {
	var text strings.Builder
	for i := range m.Height {
		for j := range m.Width {
			switch *m.Tiles[i][j].FinalState {
			case TileChargingPoint:
				text.WriteString("C")
			case TileDestination:
				text.WriteString("D")
			case TileEmpty:
				text.WriteString(" ")
			case TileRooftop:
				text.WriteString("R")
			case TileSpawner:
				text.WriteString("S")
			case TileWarehouse:
				text.WriteString("W")
			}
		}
		text.WriteString("\n")
	}

	return text.String()
}