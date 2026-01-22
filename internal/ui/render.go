package ui

import (
	"image/color"
	"swarm-drones-delivery/internal/agents/drone"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/ui/hud"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Update() error {
	g.Control()
	g.Hud.Update()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	envMap := g.Sim.Env.World()
	// adapt the window size based on the number of cells
	return envMap.Width * constants.CELL_SIZE, envMap.Height * constants.CELL_SIZE
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawMap(screen)
	g.drawObjects(screen)

	if !g.isDebugMode {
		g.drawLinesBetweenAgents(screen)
		g.drawLinesBetweenAgentAndTarget(screen)
	}

	g.drawAgents(screen)
	g.drawHUD(screen)
}

func (g *Game) drawHUD(screen *ebiten.Image) {
	if g.Hud.HudBg == nil || g.Hud.Hidden {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.Hud.PaddingX), float64(g.Hud.PaddingY))
	screen.DrawImage(g.Hud.HudBg, op)

	y := g.Hud.PaddingX + g.Hud.PaddingY + hud.FONT.Metrics().Height.Ceil()
	for _, line := range g.Hud.Lines {
		text.Draw(screen, line, hud.FONT, g.Hud.PaddingY+g.Hud.PaddingX, y, color.White)
		y += hud.FONT.Metrics().Height.Ceil()
	}

	targetX, targetY := g.mapToDrawCoords(g.Hud.TargetPosition.X, g.Hud.TargetPosition.Y)
	drawImageAt(screen, targetImg, targetX, targetY, nil)
}

func (g *Game) drawMap(screen *ebiten.Image) {
	envMap := g.Sim.Env.World()

	for y := 0.0; y < float64(envMap.Height); y++ {
		for x := 0.0; x < float64(envMap.Width); x++ {
			drawX, drawY := g.mapToDrawCoords(x, y)
			drawImageAt(screen, groundImg, drawX, drawY, WHITE)
		}
	}

	g.drawBlock(screen, envMap.Rooftops, BLACK)
	g.drawBlock(screen, envMap.Spawners, RED)
	g.drawBlock(screen, envMap.DeliveryDest, MAGENTA)
	g.drawBlock(screen, envMap.Warehouses, BLUE)
	g.drawBlock(screen, envMap.ChargingPoints, YELLOW)
}

func (g *Game) drawObjects(screen *ebiten.Image) {
	missions := g.Sim.Env.Missions()
	for _, m := range missions {
		currentPos := m.TargetDelivery.Position()
		cObjX, cObjY := g.mapToDrawCoords(currentPos.X, currentPos.Y)
		drawImageAt(screen, deliveryImg, cObjX, cObjY, MAGENTA)

		targetPos := m.Destination
		tObjX, tObjY := g.mapToDrawCoords(targetPos.X, targetPos.Y)
		drawImageAt(screen, deliveryImg, tObjX, tObjY, RED)
	}
}

func (g *Game) drawAgents(screen *ebiten.Image) {
	g.forEachSpawnedAgents(func(agt core.IAgent) {
		// If the drone is transporting a delivery, draw the delivery
		if drone, ok := agt.(*drone.Drone); ok && drone.Mission() != nil && drone.Mission().TargetDelivery != nil {
			pos := drone.Mission().TargetDelivery.Position()
			objX, objY := g.mapToDrawCoords(pos.X, pos.Y)
			drawImageAt(screen, deliveryImg, objX, objY, MAGENTA)
		}

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
