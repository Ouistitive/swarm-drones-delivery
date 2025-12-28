package Map

import (
	"os"
	"strings"
)

type Map struct {
	Width, Height int
	Cells [][]rune
}

func NewMap(width, height int) *Map {
	m := &Map{
		Width: width,
		Height: height,
	}

	m.Cells = make([][]rune, height)
	for i := range height {
		m.Cells[i] = make([]rune, width)
	}

	return m
}

func ReadMap(mapPath string) (*Map, error) {
    content, err := os.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}

	return loadMap(string(content)), nil
}

func loadMap(content string) *Map {
	lines := strings.Split(content, "\n")
	cells := make([][]rune, 0, len(lines))

	for _, line := range lines {
		row := []rune(line)
		cells = append(cells, row)
	}

	return &Map{Width: len(lines[0]), Height: len(lines), Cells: cells}
}