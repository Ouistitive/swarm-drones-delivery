package ui

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	groundImg            *ebiten.Image
	droneImg   	         *ebiten.Image
)

func init() {
	var err error
	
	groundImg, _, err = ebitenutil.NewImageFromFile("assets/ground.png")
	if err != nil {
		log.Printf("Warning: Could not load ground.png: %v", err)
	}

	droneImg, _, err = ebitenutil.NewImageFromFile("assets/drone.png")
	if err != nil {
		log.Printf("Warning: Could not load drone.png: %v", err)
	}
}
