package ui

import (
	"log"
	"os"
	"swarm-drones-delivery/internal/ui/hud"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	groundImg            *ebiten.Image
	droneImg   	         *ebiten.Image
	deliveryImg 		 *ebiten.Image
	targetImg 			 *ebiten.Image
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

	deliveryImg, _, err = ebitenutil.NewImageFromFile("assets/delivery.png")
	if err != nil {
		log.Printf("Warning: Could not load delivery.png: %v", err)
	}

	targetImg, _, err = ebitenutil.NewImageFromFile("assets/target.png")
	if err != nil {
		log.Printf("Warning: Could not load target.png: %v", err)
	}

	fontBytes, err := os.ReadFile("assets/fonts/Monaco.ttf")
	if err != nil {
		panic(err)
	}

	ttf, err := opentype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	hud.FONT, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    16,
		DPI:     96,
		Hinting: font.HintingFull,
	})
}
