package ui

import (
	"image/color"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) mapToDrawCoords(mapX float64, mapY float64) (float64, float64) {
	return mapX * float64(constants.CELL_SIZE), mapY * float64(constants.CELL_SIZE)
}

func (g *Game) mapToDrawCoordsCentered(mapX float64, mapY float64) (float64, float64) {
	return mapX*float64(constants.CELL_SIZE) + constants.HALF_CELL_SIZE, mapY*float64(constants.CELL_SIZE) + constants.HALF_CELL_SIZE
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
	options.GeoM.Translate(x, y)
	screen.DrawImage(img, options)
}

func (g *Game) forEachSpawnedAgents(f func(agt core.IAgent)) {
	agts := g.Sim.Env.SpawnedAgents()

	for _, agt := range agts {
		f(agt)
	}
}

func (g *Game) drawBlock(screen *ebiten.Image, poses []world.Position, color *ebiten.ColorScale) {
	for _, pos := range poses {
		drawX, drawY := g.mapToDrawCoords(pos.X, pos.Y)
		drawImageAt(screen, groundImg, drawX, drawY, color)
	}
}

func initVisionCircle(visionRadiusPx int) {
    diameter := visionRadiusPx * constants.CELL_SIZE
    visionCircle = ebiten.NewImage(diameter, diameter)

    vector.FillCircle(
        visionCircle,
        float32(visionRadiusPx),
        float32(visionRadiusPx),
        float32(visionRadiusPx),
        color.RGBA{0, 0, 0, 10},
        false,
    )
}
