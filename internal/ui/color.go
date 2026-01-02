package ui

import "github.com/hajimehoshi/ebiten/v2"

var (
	BLACK *ebiten.ColorScale
	WHITE *ebiten.ColorScale
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
}
