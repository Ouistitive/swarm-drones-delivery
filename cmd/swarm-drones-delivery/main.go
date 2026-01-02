package main

import (
	"log"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Swarm drones delivery simulation")
	game := ui.NewGame("maps/layout.txt")
	
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}