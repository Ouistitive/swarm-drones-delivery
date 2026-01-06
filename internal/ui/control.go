package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Control() {
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.isDebugMode = !g.isDebugMode
	}
}