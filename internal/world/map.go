package world

import "math/rand/v2"

type Position struct {
	X, Y float64
}

type Map struct {
	Width, Height int
	Cells [][]rune
	Walls []Position
}

func NewPosition(x, y float64) Position {
	return Position{ X: x, Y: y }
}

func (m *Map) RandomPosition() Position {
	return NewPosition(
		rand.Float64() * float64(m.Width-1), 
		rand.Float64() * float64(m.Height-1),
	)
}

func NewMap(width, height int) *Map {
	m := &Map{
		Width: width,
		Height: height,
		Walls: make([]Position, 0),
	}

	m.Cells = make([][]rune, height)
	for i := range height {
		m.Cells[i] = make([]rune, width)
	}

	return m
}