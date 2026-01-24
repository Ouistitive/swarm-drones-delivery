package ui

import (
	"image/color"
	"log"
	"swarm-drones-delivery/internal/constants"
	"swarm-drones-delivery/internal/simulation"
	"swarm-drones-delivery/internal/ui/hud"
	"swarm-drones-delivery/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Sim         *simulation.Simulation
	Hud         hud.Hud
	isDebugMode bool

	groundLayer       *ebiten.Image
	rooftopsLayer     *ebiten.Image
	spawnersLayer     *ebiten.Image
	deliveryDestLayer *ebiten.Image
	warehousesLayer   *ebiten.Image
	chargingLayer     *ebiten.Image

	visionCircle *ebiten.Image
}

func NewGame(mapPath, addressesPath string) *Game {
	m, err := world.ReadMap(mapPath, addressesPath)
	if err != nil {
		log.Fatal("Cannot load map")
	}
	sim := simulation.NewSimulation(m)

	g := &Game{
		Sim:         sim,
		isDebugMode: false,
	}
	g.buildGroundCache()
	g.buildVisionCircle()

	return g
}

func (g *Game) buildGroundCache() {
	envMap := g.Sim.Env.World()

	w := envMap.Width * constants.CELL_SIZE
	h := envMap.Height * constants.CELL_SIZE

	g.groundLayer = ebiten.NewImage(w, h)
	g.rooftopsLayer = ebiten.NewImage(w, h)
	g.spawnersLayer = ebiten.NewImage(w, h)
	g.deliveryDestLayer = ebiten.NewImage(w, h)
	g.warehousesLayer = ebiten.NewImage(w, h)
	g.chargingLayer = ebiten.NewImage(w, h)

	for y := 0; y < envMap.Height; y++ {
		for x := 0; x < envMap.Width; x++ {
			drawX := float64(x * constants.CELL_SIZE)
			drawY := float64(y * constants.CELL_SIZE)
			drawImageAt(g.groundLayer, groundImg, drawX, drawY, WHITE)
		}
	}

	g.drawBlock(g.rooftopsLayer, envMap.Rooftops, BLACK)
	g.drawBlock(g.spawnersLayer, envMap.Spawners, RED)
	g.drawBlock(g.warehousesLayer, envMap.Warehouses, BLUE)

	for _, delDest := range envMap.DeliveryDests {
		drawX, drawY := g.mapToDrawCoords(delDest.Pos.X, delDest.Pos.Y)
		drawImageAt(g.deliveryDestLayer, groundImg, drawX, drawY, YELLOW)
	}

	for _, chargingPtn := range envMap.ChargingPoints {
		drawX, drawY := g.mapToDrawCoords(chargingPtn.Pos.X, chargingPtn.Pos.Y)
		drawImageAt(g.chargingLayer, groundImg, drawX, drawY, YELLOW)
	}
}

func (g *Game) buildVisionCircle() {
	radius := float32(constants.VISION_RANGE) * float32(constants.CELL_SIZE)
	diameter := int(radius * 2)

	g.visionCircle = ebiten.NewImage(diameter, diameter)
	g.visionCircle.Fill(color.RGBA{0, 0, 0, 0})

	vector.FillCircle(g.visionCircle, radius, radius, radius, color.RGBA{0, 0, 0, 10}, false)
}
