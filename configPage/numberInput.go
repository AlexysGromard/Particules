package configPage

import (
	"image"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

var (
	touchIsPressed bool = false
)

// Création de la structure TextInput
type NumberInput struct {
	x, y, width, height int
	imageNormal         *ebiten.Image
	imageHover          *ebiten.Image
	imageDisabled       *ebiten.Image
	hover               bool
	pressed             bool
	disabled            bool
	number              *int
	fontface            font.Face
	Text                *Text
}

// Draw
func (t *NumberInput) draw(screen *ebiten.Image) {
	var img *ebiten.Image
	if t.disabled {
		img = t.imageDisabled
	} else if t.hover || t.pressed {
		img = t.imageHover
	} else {
		img = t.imageNormal
	}
	// Ajouter le scroll
	// Draw text
	t.Text.y = t.y + t.height/2 + 7
	t.Text.draw(screen)
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(t.x), float64(t.y))
	screen.DrawImage(nineSliceInput(img, t.width, t.height), opt)
}

// Update
func (t *NumberInput) update(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	t.hover = false
	// Si la souris est sur le bouton
	if (x > t.x && x < t.x+t.width && y > t.y && y < t.y+t.height) && !t.disabled {
		t.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			t.pressed = true
			t.Text.color = color.RGBA{226, 199, 90, 255}
		}
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && (x < t.x || x > t.x+t.width || y < t.y || y > t.y+t.height) {
		t.pressed = false
		t.Text.color = color.RGBA{127, 139, 148, 255}
	}
	// Verifie si aucune touche n'est appuyée
	if !ebiten.IsKeyPressed(ebiten.Key0) && !ebiten.IsKeyPressed(ebiten.KeyNumpad0) && !ebiten.IsKeyPressed(ebiten.Key1) && !ebiten.IsKeyPressed(ebiten.KeyNumpad1) && !ebiten.IsKeyPressed(ebiten.Key2) && !ebiten.IsKeyPressed(ebiten.KeyNumpad2) && !ebiten.IsKeyPressed(ebiten.Key3) && !ebiten.IsKeyPressed(ebiten.KeyNumpad3) && !ebiten.IsKeyPressed(ebiten.Key4) && !ebiten.IsKeyPressed(ebiten.KeyNumpad4) && !ebiten.IsKeyPressed(ebiten.Key5) && !ebiten.IsKeyPressed(ebiten.KeyNumpad5) && !ebiten.IsKeyPressed(ebiten.Key6) && !ebiten.IsKeyPressed(ebiten.KeyNumpad6) && !ebiten.IsKeyPressed(ebiten.Key7) && !ebiten.IsKeyPressed(ebiten.KeyNumpad7) && !ebiten.IsKeyPressed(ebiten.Key8) && !ebiten.IsKeyPressed(ebiten.KeyNumpad8) && !ebiten.IsKeyPressed(ebiten.Key9) && !ebiten.IsKeyPressed(ebiten.KeyNumpad9) && !ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		touchIsPressed = false
	}
	// Si t.pressed est vrai, on peut écrire dans le textInput
	if t.pressed && !touchIsPressed {
		var str string = strconv.Itoa(*t.number)
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
		*t.number, _ = strconv.Atoi(str)
	}
	// Update text
	t.Text.text = strconv.Itoa(*t.number)
	t.Text.update(screen)
	// Ajouter le scroll
	t.y += ScrollY
	// Draw le numberInput
	t.draw(screen)
}

// Création d'un nouveau TextInput
func newTextInput(x, y, width, height int, images []*ebiten.Image, disabled bool, number *int, fontFace font.Face) *NumberInput {
	return &NumberInput{
		x:             x,
		y:             y,
		width:         width,
		height:        height,
		imageNormal:   images[0],
		imageHover:    images[1],
		imageDisabled: images[2],
		disabled:      disabled,
		number:        number,
		fontface:      fontFace,
		Text:          newText(x+10, y+height/2+7, strconv.Itoa(*number), fontFace, color.RGBA{127, 139, 148, 255}),
	}
}

// Fonction de NineSlice
// Cette fonction permet de découper une image en 9 parties pour pouvoir l'étirer sans que les bords ne se déforment
// Entrées : image (image à découper), width, height (taille de l'image)
func nineSliceInput(img *ebiten.Image, width, height int) *ebiten.Image {
	w, h := img.Size()
	// On découpe l'image en 9 parties
	imgTab := make([]*ebiten.Image, 3)
	// imgTab[0] = Gauche
	imgTab[0] = img.SubImage(image.Rect(0, 0, 9, h)).(*ebiten.Image)
	// imgTab[1] = Millieu
	imgTab[1] = img.SubImage(image.Rect(9, 0, w-2, h)).(*ebiten.Image)
	// imgTab[2] = coin haut droit
	imgTab[2] = img.SubImage(image.Rect(w-2, 0, w, h)).(*ebiten.Image)
	// ON créée une nouvelle image
	imgOut := ebiten.NewImage(width, height)
	// On dessine les 9 parties de l'image sur la nouvelle image
	opt := &ebiten.DrawImageOptions{}
	// Mettre les 4 coins
	// Partie  à gauche
	opt.GeoM.Scale(1, 1)
	opt.GeoM.Translate(0, 0)
	imgOut.DrawImage(imgTab[0], opt)
	// Partie en haut à droite
	opt.GeoM.Translate(float64(width)-23, 0)
	imgOut.DrawImage(imgTab[2], opt)
	// Partie au milieu
	opt.GeoM.Scale(3.85, 1)
	opt.GeoM.Translate(float64(-w*8-55), 0)
	imgOut.DrawImage(imgTab[1], opt)
	return imgOut
}
