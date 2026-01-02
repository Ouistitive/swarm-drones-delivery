package ui

import (
	"swarm-drones-delivery/internal/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	envMap := g.Sim.Env.Map
	// adapt the window size based on the number of cells
	return envMap.Width * constants.CELL_SIZE, envMap.Height * constants.CELL_SIZE
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawMap(screen)
}

func (g *Game) DrawMap(screen *ebiten.Image) {
	envMap := g.Sim.Env.Map
	for y := 0.0; y < float64(envMap.Height); y++ {
		for x := 0.0; x < float64(envMap.Width); x++ {
			drawX, drawY := g.mapToDrawCoords(x, y)
			drawImageAt(screen, groundImg, drawX, drawY, WHITE)
		}
	}

	for _, pos := range envMap.Walls {
		drawX, drawY := g.mapToDrawCoords(pos.X, pos.Y)
		drawImageAt(screen, groundImg, drawX, drawY, BLACK)
	}
}

func (g *Game) mapToDrawCoords(mapX float64, mapY float64) (float64, float64) {
	return mapX * float64(constants.CELL_SIZE), mapY * float64(constants.CELL_SIZE)
}

func drawImageAt(screen *ebiten.Image, img *ebiten.Image, x, y float64, colorScale *ebiten.ColorScale) {
	if img == nil {
		return
	}
	options := &ebiten.DrawImageOptions{}
	if colorScale != nil {
		options.ColorScale = *colorScale
	}

	options.GeoM.Scale(float64(constants.CELL_SIZE)/float64(img.Bounds().Dx()), float64(constants.CELL_SIZE)/float64(img.Bounds().Dy()))
	options.GeoM.Translate(float64(x), float64(y))

	screen.DrawImage(img, options)
}
