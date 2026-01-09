package main

import (
	"log"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowSize(constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Swarm drones delivery simulation")

	game := ui.NewGame(constants.LAYOUT_PATH)
	game.Sim.Run()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}