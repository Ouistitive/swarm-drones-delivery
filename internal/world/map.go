package world

import (
	"math"
	"math/rand"
)


type Map struct {
	Width, Height int
	Cells         [][]rune
	Walls         []Position
	Spawners      []Position
}

func NewMap(width, height int) *Map {
	m := &Map{
		Width:    width,
		Height:   height,
		Walls:    make([]Position, 0),
		Spawners: make([]Position, 0),
	}

	m.Cells = make([][]rune, height)
	for i := range height {
		m.Cells[i] = make([]rune, width)
	}

	return m
}

func (m *Map) RandomPosition() Position {
	return NewPosition(
		math.Round(rand.Float64()*float64(m.Width-1)),
		math.Round(rand.Float64()*float64(m.Height-1)),
	)
}

func (m *Map) RandomSpawner() (Position, int) {
	n := rand.Intn(len(m.Spawners))
    return m.Spawners[n], n
}
