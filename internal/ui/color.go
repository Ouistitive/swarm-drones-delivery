package ui

import "github.com/hajimehoshi/ebiten/v2"

var (
	BLACK *ebiten.ColorScale
	WHITE *ebiten.ColorScale

	RED     *ebiten.ColorScale
	MAGENTA *ebiten.ColorScale
	BLUE	*ebiten.ColorScale
	YELLOW	*ebiten.ColorScale
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

	MAGENTA = &ebiten.ColorScale{}
	MAGENTA.SetR(255)
	MAGENTA.SetG(0)
	MAGENTA.SetB(255)
	
	BLUE = &ebiten.ColorScale{}
	BLUE.SetR(0)
	BLUE.SetG(0)
	BLUE.SetB(255)

	YELLOW = &ebiten.ColorScale{}
	YELLOW.SetR(255)
	YELLOW.SetG(255)
	YELLOW.SetB(0)
}
