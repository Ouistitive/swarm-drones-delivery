package ui

import "github.com/hajimehoshi/ebiten/v2"

var (
	BLACK 	*ebiten.ColorScale
	WHITE 	*ebiten.ColorScale

	RED 	*ebiten.ColorScale
)

func init() {
	BLACK = &ebiten.ColorScale{}
	BLACK.SetR(0)
	BLACK.SetG(0)
	BLACK.SetB(0)

	WHITE = &ebiten.ColorScale{}
	WHITE.SetR(255)
	WHITE.SetG(255)
	WHITE.SetB(255)

	RED = &ebiten.ColorScale{}
	RED.SetR(255)
	RED.SetG(0)
	RED.SetB(0)
}
