package configPage

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Value struct {
	x, y     int
	value    *float64
	fontFace font.Face
	color    color.Color
}

// Draw dessine le texte
func (v *Value) Draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("%v", *v.value), v.fontFace, v.x, v.y, v.color)
}

// Update met à jour le texte
func (v *Value) Update(screen *ebiten.Image) {
	// Mettre à jour le texte
	// Value = 2 chiffres après la virgule
	*v.value = float64(int(*v.value*100)) / 100
	text.Draw(screen, fmt.Sprintf("%v", *v.value), v.fontFace, v.x, v.y, v.color)
}

// NewText crée un nouveau texte
func newValue(x, y int, value *float64, fontFace font.Face, color color.Color) *Value {
	return &Value{
		x:        x,
		y:        y,
		value:    value,
		fontFace: fontFace,
		color:    color,
	}
}
