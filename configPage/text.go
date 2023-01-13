package configPage

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Text struct {
	x, y     int
	text     string
	fontFace font.Face
	color    color.Color
}

// Draw dessine le texte
func (t *Text) Draw(screen *ebiten.Image) {
	text.Draw(screen, t.text, t.fontFace, t.x, t.y, t.color)
}

func (t *Text) updateText(screen *ebiten.Image) {
	// Mettre à jour le texte
	text.Draw(screen, t.text, t.fontFace, t.x, t.y, t.color)
}

// NewText crée un nouveau texte
func newText(x, y int, text string, fontFace font.Face, color color.Color) *Text {
	return &Text{
		x:        x,
		y:        y,
		text:     text,
		fontFace: fontFace,
		color:    color,
	}
}
