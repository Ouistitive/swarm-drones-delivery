package world

type Position struct {
	X, Y float64
}

func NewPosition(x, y float64) Position {
	return Position{ X: x, Y: y }
}

func NullPosition() Position {
	return NewPosition(0.0, 0.0)
}
