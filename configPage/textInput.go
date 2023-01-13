package configPage

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

var (
	touchIsPressed bool = false
)

// Création de la structure TextInput
type TextInput struct {
	x, y, width, height int
	imageNormal         *ebiten.Image
	imageHover          *ebiten.Image
	imagePressed        *ebiten.Image
	hover               bool
	pressed             bool
	text                *int
	fontface            font.Face
	color               color.Color
	Text                *Text
}

// Draw
func (t *TextInput) Draw(screen *ebiten.Image) {
	var img *ebiten.Image
	if t.hover || t.pressed {
		img = t.imageHover
	} else {
		img = t.imageNormal
	}
	// Draw text
	t.Text.Draw(screen)
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(t.x), float64(t.y))
	screen.DrawImage(img, opt)
}

// Update
func (t *TextInput) updateTextInput(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	t.hover = false
	// Si la souris est sur le bouton
	if x > t.x && x < t.x+t.width && y > t.y && y < t.y+t.height {
		t.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			t.pressed = true
		}
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && (x < t.x || x > t.x+t.width || y < t.y || y > t.y+t.height) {
		t.pressed = false
	}
	// Verifie si aucune touche n'est appuyée
	if !ebiten.IsKeyPressed(ebiten.Key0) && !ebiten.IsKeyPressed(ebiten.KeyNumpad0) && !ebiten.IsKeyPressed(ebiten.Key1) && !ebiten.IsKeyPressed(ebiten.KeyNumpad1) && !ebiten.IsKeyPressed(ebiten.Key2) && !ebiten.IsKeyPressed(ebiten.KeyNumpad2) && !ebiten.IsKeyPressed(ebiten.Key3) && !ebiten.IsKeyPressed(ebiten.KeyNumpad3) && !ebiten.IsKeyPressed(ebiten.Key4) && !ebiten.IsKeyPressed(ebiten.KeyNumpad4) && !ebiten.IsKeyPressed(ebiten.Key5) && !ebiten.IsKeyPressed(ebiten.KeyNumpad5) && !ebiten.IsKeyPressed(ebiten.Key6) && !ebiten.IsKeyPressed(ebiten.KeyNumpad6) && !ebiten.IsKeyPressed(ebiten.Key7) && !ebiten.IsKeyPressed(ebiten.KeyNumpad7) && !ebiten.IsKeyPressed(ebiten.Key8) && !ebiten.IsKeyPressed(ebiten.KeyNumpad8) && !ebiten.IsKeyPressed(ebiten.Key9) && !ebiten.IsKeyPressed(ebiten.KeyNumpad9) && !ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		touchIsPressed = false
	}
	// Si t.pressed est vrai, on peut écrire dans le textInput
	if t.pressed && !touchIsPressed {
		var str string = strconv.Itoa(*t.text)
		if ebiten.IsKeyPressed(ebiten.Key0) || ebiten.IsKeyPressed(ebiten.KeyNumpad0) {
			str += "0"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key1) || ebiten.IsKeyPressed(ebiten.KeyNumpad1) {
			str += "1"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key2) || ebiten.IsKeyPressed(ebiten.KeyNumpad2) {
			str += "2"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key3) || ebiten.IsKeyPressed(ebiten.KeyNumpad3) {
			str += "3"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key4) || ebiten.IsKeyPressed(ebiten.KeyNumpad4) {
			str += "4"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key5) || ebiten.IsKeyPressed(ebiten.KeyNumpad5) {
			str += "5"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key6) || ebiten.IsKeyPressed(ebiten.KeyNumpad6) {
			str += "6"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key7) || ebiten.IsKeyPressed(ebiten.KeyNumpad7) {
			str += "7"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key8) || ebiten.IsKeyPressed(ebiten.KeyNumpad8) {
			str += "8"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.Key9) || ebiten.IsKeyPressed(ebiten.KeyNumpad9) {
			str += "9"
			touchIsPressed = true
		} else if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
			if len(str) > 0 {
				str = str[:len(str)-1]
				touchIsPressed = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			t.pressed = false
		}
		*t.text, _ = strconv.Atoi(str)
	}
	// Update text
	t.Text.text = strconv.Itoa(*t.text)
	t.Text.updateText(screen)
}

// Création d'un nouveau TextInput
func newTextInput(x, y, width, height int, imageNormal, imageHover, imagePressed *ebiten.Image, text *int, fontFace font.Face, color color.Color) *TextInput {
	return &TextInput{
		x:            x,
		y:            y,
		width:        width,
		height:       height,
		imageNormal:  imageNormal,
		imageHover:   imageHover,
		imagePressed: imagePressed,
		text:         text,
		fontface:     fontFace,
		color:        color,
		Text:         newText(x+10, y+height/2-3, strconv.Itoa(*text), fontFace, color),
	}
}
