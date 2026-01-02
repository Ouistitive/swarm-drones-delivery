package world

type Position struct {
	x, y float64
}

type Map struct {
	Width, Height int
	Cells [][]rune
	Walls []Position
}

func NewPosition(x, y float64) Position {
	return Position{ x: x, y: y }
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