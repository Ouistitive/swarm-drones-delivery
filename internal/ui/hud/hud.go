package hud

import (
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

var FONT font.Face

type Hud struct {
	TargetPosition 		world.Position
	PaddingX, PaddingY  int
	HudWidth, HudHeight int
	HudBg               *ebiten.Image
	Lines               []string
	Hidden 				bool
	
	selectedAgt 		core.ClickableEntity
}

func NewHud() *Hud {
	return &Hud{
		TargetPosition:  world.NewPosition(10, 10),
		PaddingX:        10,
		PaddingY:        5,
		Hidden:          true,
	}
}