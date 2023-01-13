package configPage

import (
	"image/color"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Création de la structure Button
type Button struct {
	x, y, width, height int
	imageNormal         *ebiten.Image
	imagePressed        *ebiten.Image
	imageHover          *ebiten.Image
	text                string
	textX, textY        int
	font                font.Face
	pressed             bool
	hover               bool
	onClick             func()
}

// Impression des bouttons
// Cette fonction met à jour l'affichage des bouttons en fonction de leur état (hover, pressed)
func (b *Button) Draw(screen *ebiten.Image) {
	var img *ebiten.Image
	if b.pressed {
		img = b.imagePressed
	} else if b.hover {
		img = b.imageHover
	} else {
		img = b.imageNormal
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(b.x), float64(b.y))

	screen.DrawImage(img, opt)
	text.Draw(screen, b.text, b.font, b.textX, b.textY, color.White)
}

// Mise à jour des bouttons
// Cette fonction regarde si la souris est sur un boutton et si elle est enfoncée
// Elle regarde vérifie si le bouton en bas à droite est toujours à la bonne position (en fonction de l'affichage de la fenêtre)
func (b *Button) updateButton(screen *ebiten.Image) {
	// Position de la souris
	x, y := ebiten.CursorPosition()
	b.hover = false
	// Si la souris est sur le boutton
	if x > b.x && x < b.x+b.width && y > b.y && y < b.y+b.height {
		b.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			b.pressed = true
		}
	} else {
		b.pressed = false
	}
	// Si le bouton en bas à gauche n'est plus à la bonne place en fonction de la taille d'écran
	if b.x < config.General.WindowSizeX-b.width || b.x > config.General.WindowSizeX-b.width+10 || b.y < config.General.WindowSizeY-b.height || b.y > config.General.WindowSizeY-b.height+10 {
		b.x = config.General.WindowSizeX - b.width - 10
		b.y = config.General.WindowSizeY - b.height - 10
		b.textX = b.x + 10
		b.textY = b.y + b.height/2
	}
}

// NewButton crée un nouveau boutton en fonction des paramètres
// Entrées : x,y (position), width, height (taille), images (liste des images), text (texte du boutton), font (police d'écriture), onClick (fonction à appeler quand on clique sur le boutton)
func newButton(x, y, width, height int, images []*ebiten.Image, text string, font font.Face, onClick func()) *Button {
	return &Button{
		x:            x,
		y:            y,
		width:        width,
		height:       height,
		imageNormal:  images[0],
		imagePressed: images[1],
		imageHover:   images[2],
		text:         text,
		textX:        x,
		textY:        y,
		font:         font,
		onClick:      onClick,
	}
}
