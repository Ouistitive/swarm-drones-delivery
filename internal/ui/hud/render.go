package hud

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (h *Hud) Update() {
	if h.selectedAgt != nil {
		h.TargetPosition = h.selectedAgt.Position()
		// msg := h.getAgentSelectionMessage()
		msg := "coucou"
		h.prepareRender(msg)
	}
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
	h.HudWidth = maxWidth + h.PaddingX*2
	h.HudHeight = len(lines)*lineHeight + h.PaddingY*2
	h.HudBg = ebiten.NewImage(h.HudWidth, h.HudHeight)
	h.HudBg.Fill(color.RGBA{0, 0, 0, 180})
}
