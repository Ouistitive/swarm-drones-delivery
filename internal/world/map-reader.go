package world

import (
	"os"
	"strings"
)

func ReadMap(mapPath string) (*Map, error) {
	content, err := os.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}

	return loadMap(string(content)), nil
}

func loadMap(content string) *Map {
	lines := strings.Split(content, "\n")
	cells := make([][]rune, 0)
	walls := make([]Position, 0)
	spawners := make([]Position, 0)

	y := 0.0
	for _, line := range lines {
		row := []rune(line)
		cells = append(cells, row)
		for x, r := range line {
			switch r {
			case 'W':
				walls = append(walls, NewPosition(float64(x), y))
			case 'S':
				spawners = append(spawners, NewPosition(float64(x), y))
			}
		}

		y++
	}

	return &Map{Width: len(lines[0]), Height: len(lines), Cells: cells, Walls: walls, Spawners: spawners}
}
