package ui

import (
	"fmt"
	"math"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Control() {
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.isDebugMode = !g.isDebugMode
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.handleMouseClick()
	}
}

func (g *Game) handleMouseClick() {
	mouseX, mouseY := ebiten.CursorPosition()

	// convert screen to map coordinates
	mapX := float64((mouseX) / constants.CELL_SIZE)
	mapY := float64((mouseY) / constants.CELL_SIZE)

	agt := g.isMouseClickedOnAgent(mapX, mapY)
	if agt != nil {
		fmt.Println(agt.ID())
		g.Hud.SetAgent(agt)
	}
}

func (g *Game) isMouseClickedOnAgent(mapX, mapY float64) core.IAgent {
	for _, a := range g.Sim.Env.Agents() {
		dx := math.Abs(a.Position().X - mapX)
		dy := math.Abs(a.Position().Y - mapY)

		if dx < 0.5 && dy < 0.5 {
			return a
		}
	}

	return nil
}