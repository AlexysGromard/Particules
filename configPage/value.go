package configPage

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type ValueF struct {
	x, y     int
	value    *float64
	fontFace font.Face
	color    color.Color
}

type ValueI struct {
	x, y     int
	value    *int
	fontFace font.Face
	color    color.Color
}

// Draw dessine le texte
func (v *ValueF) draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("%v", *v.value), v.fontFace, v.x, v.y, v.color)
}
func (v *ValueI) draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("%v", *v.value), v.fontFace, v.x, v.y, v.color)
}

// Update met à jour le texte
func (v *ValueF) update(screen *ebiten.Image) {
	// Mettre à jour le texte
	// Value = 2 chiffres après la virgule
	*v.value = float64(int(*v.value*100)) / 100
	// Ajouter le scroll
	v.y += ScrollY
	// Draw le texte
	v.draw(screen)
}
func (v *ValueI) update(screen *ebiten.Image) {
	// Mettre à jour le texte
	// Value = 2 chiffres après la virgule
	*v.value = *v.value * 100 / 100
	// Ajouter le scroll
	v.y += ScrollY
	// Draw le texte
	v.draw(screen)
}

// NewText crée un nouveau texte
func newValueFloat(x, y int, value *float64, fontFace font.Face, color color.Color) *ValueF {
	return &ValueF{
		x:        x,
		y:        y,
		value:    value,
		fontFace: fontFace,
		color:    color,
	}
}

func newValueInt(x, y int, value *int, fontFace font.Face, color color.Color) *ValueI {
	return &ValueI{
		x:        x,
		y:        y,
		value:    value,
		fontFace: fontFace,
		color:    color,
	}
}
