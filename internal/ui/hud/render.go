package hud

import (
	"image/color"
	"strings"
	"swarm-drones-delivery/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (h *Hud) Update() {
	if h.selectedAgt != nil {
		h.TargetPosition = h.selectedAgt.Position()
		h.prepareRender(h.selectedAgt.GetDisplayData())
	}
}

func (h *Hud) SetAgent(agt core.ClickableEntity) {
	h.selectedAgt = agt
}

// Determine the width and height the background based on the text
func (h *Hud) prepareRender(msg string) {
	lines := strings.Split(msg, "\n")
	h.Lines = lines
	lineHeight := FONT.Metrics().Height.Ceil()
	maxWidth := 0
	for _, line := range lines {
		bounds := text.BoundString(FONT, line)
		width := bounds.Max.X + 1
		if width > maxWidth {
			maxWidth = width
		}
	}
	h.HudWidth = maxWidth + h.PaddingX * 2
	h.HudHeight = len(lines) * lineHeight + h.PaddingY * 2
	h.HudBg = ebiten.NewImage(h.HudWidth, h.HudHeight)
	h.HudBg.Fill(color.RGBA{0, 0, 0, 180})
}
