package hud

import (
	"swarm-drones-delivery/internal/core"
	"swarm-drones-delivery/internal/world"

	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten/v2"
)

var FONT font.Face

type Hud struct {
	TargetPosition 					 world.Position
	PaddingX, PaddingY               int
	HudWidth, HudHeight              int
	HudBg                            *ebiten.Image
	Lines                            []string
	
	selectedAgt core.IAgent
}

func NewHud() *Hud {
	return &Hud{
		TargetPosition:  world.NewPosition(10, 10),
		PaddingX:        10,
		PaddingY:        5,
		// hidden:          true,
		// DisplayAgentPaths: false,
	}
}

// func (h *Hud) Hidden() bool {
// 	return h.hidden
// }

// func (h *Hud) ToggleHidden() {
// 	h.hidden = !h.hidden
// }