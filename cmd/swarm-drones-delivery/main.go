package main

import (
	"log"
	"os"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("Error loading .env file: %s", err)
	}

	LAYOUT_PATH := os.Getenv("LAYOUT_PATH")
	
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowSize(constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Swarm drones delivery simulation")

	game := ui.NewGame(LAYOUT_PATH)
	game.Sim.Run()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}