package ui

import (
	"image/color"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Update() error {
	g.Control()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	envMap := g.Sim.Env.World()
	// adapt the window size based on the number of cells
	return envMap.Width * constants.CELL_SIZE, envMap.Height * constants.CELL_SIZE
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawMap(screen)
	if !g.isDebugMode {
		g.drawLinesBetweenAgents(screen)
		g.drawLinesBetweenAgentAndTarget(screen)
	}
	g.drawAgents(screen)
}

func (g *Game) drawMap(screen *ebiten.Image) {
	envMap := g.Sim.Env.World()

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

	for _, pos := range envMap.Spawners {
		drawX, drawY := g.mapToDrawCoords(pos.X, pos.Y)
		drawImageAt(screen, groundImg, drawX, drawY, RED)
	}
}

func (g *Game) drawAgents(screen *ebiten.Image) {
	g.forEachSpawnedAgents(func(agt core.IAgent) {
		targX, targY := g.mapToDrawCoords(agt.TargetPos().X, agt.TargetPos().Y)
		drawImageAt(screen, droneImg, targX, targY, RED)
		agtX, agtY := g.mapToDrawCoords(agt.Position().X, agt.Position().Y)
		drawImageAt(screen, droneImg, agtX, agtY, nil)
	})
}

func (g *Game) drawLinesBetweenAgents(screen *ebiten.Image) {
	g.forEachSpawnedAgents(func(agt core.IAgent) {
		drawX, drawY := g.mapToDrawCoordsCentered(agt.Position().X, agt.Position().Y)
		for _, surrAgt := range agt.SurroundingAgents() {
			surrAgtX, surrAgtY := g.mapToDrawCoordsCentered(surrAgt.Position().X, surrAgt.Position().Y)
			vector.StrokeLine(screen, float32(drawX), float32(drawY), float32(surrAgtX), float32(surrAgtY), 2, color.RGBA{0, 100, 255, 255}, false)
		}
	})
}

func (g *Game) drawLinesBetweenAgentAndTarget(screen *ebiten.Image) {
	g.forEachSpawnedAgents(func(agt core.IAgent) {
		drawX, drawY := g.mapToDrawCoordsCentered(agt.Position().X, agt.Position().Y)
		tX, tY := g.mapToDrawCoordsCentered(agt.TargetPos().X, agt.TargetPos().Y)
		vector.StrokeLine(screen, float32(drawX), float32(drawY), float32(tX), float32(tY), 1, color.RGBA{255, 0, 0, 255}, false)
	})
}

func (g *Game) mapToDrawCoords(mapX float64, mapY float64) (float64, float64) {
	return mapX * float64(constants.CELL_SIZE), mapY * float64(constants.CELL_SIZE)
}

func (g *Game) mapToDrawCoordsCentered(mapX float64, mapY float64) (float64, float64) {
	return mapX * float64(constants.CELL_SIZE) + constants.HALF_CELL_SIZE, mapY * float64(constants.CELL_SIZE) + constants.HALF_CELL_SIZE
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