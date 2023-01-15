package configPage

import (
	"image"
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

	screen.DrawImage(nineSlice(img, b.width, b.height), opt)
	// Draw text in the center of the button
	text.Draw(screen, b.text, b.font, b.textX+33, b.textY+2, color.RGBA{185, 204, 216, 255})
}

// Mise à jour des bouttons
// Cette fonction regarde si la souris est sur un boutton et si elle est enfoncée
// Elle regarde vérifie si le bouton en bas à droite est toujours à la bonne position (en fonction de l'affichage de la fenêtre)
func (b *Button) Update(screen *ebiten.Image) {
	// Position de la souris
	x, y := ebiten.CursorPosition()
	b.hover = false
	// Si la souris est sur le boutton
	if x > b.x && x < b.x+b.width && y > b.y && y < b.y+b.height {
		b.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			b.pressed = true
			b.onClick()
		} else {
			b.pressed = false
		}
	} else {
		b.pressed = false
	}
	// Si le bouton en bas à gauche n'est plus à la bonne place en fonction de la taille d'écran
	if b.x < config.General.WindowSizeX-b.width || b.x > config.General.WindowSizeX-b.width+30 || b.y < config.General.WindowSizeY-b.height || b.y > config.General.WindowSizeY-b.height+30 {
		b.x = config.General.WindowSizeX - b.width - 30
		b.y = config.General.WindowSizeY - b.height - 30
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

// Fonction de NineSlice
// Cette fonction permet de découper une image en 9 parties pour pouvoir l'étirer sans que les bords ne se déforment
// Entrées : image (image à découper), width, height (taille de l'image)
func nineSlice(img *ebiten.Image, width, height int) *ebiten.Image {
	w, h := img.Size()
	// On découpe l'image en 9 parties
	imgTab := make([]*ebiten.Image, 5)
	// imgTab[0] = coin haut gauche
	imgTab[0] = img.SubImage(image.Rect(0, 0, 23, 21)).(*ebiten.Image)
	// imgTab[1] = centre
	imgTab[1] = img.SubImage(image.Rect(24, 0, w-23, height)).(*ebiten.Image)
	// imgTab[2] = coin haut droit
	imgTab[2] = img.SubImage(image.Rect(w-23, 0, w, 20)).(*ebiten.Image)
	// imgTab[3] = bord bas gauche
	imgTab[3] = img.SubImage(image.Rect(0, h-21, 23, h)).(*ebiten.Image)
	// imgTab[5] = coin bas droit
	imgTab[4] = img.SubImage(image.Rect(w-23, h-21, w, h)).(*ebiten.Image)
	// ON créée une nouvelle image
	imgOut := ebiten.NewImage(width, height)
	// On dessine les 9 parties de l'image sur la nouvelle image
	opt := &ebiten.DrawImageOptions{}
	// Mettre les 4 coins
	// Partie en haut à gauche
	opt.GeoM.Scale(1, 1)
	opt.GeoM.Translate(0, 0)
	imgOut.DrawImage(imgTab[0], opt)
	// Partie en haut à droite
	opt.GeoM.Translate(float64(width)-23, 0)
	imgOut.DrawImage(imgTab[2], opt)
	// Partie en bas à gauche
	opt.GeoM.Translate(-float64(width)+23, float64(height)-30)
	imgOut.DrawImage(imgTab[3], opt)
	// Partie en bas à droite
	opt.GeoM.Translate(float64(width)-23, 0)
	imgOut.DrawImage(imgTab[4], opt)
	// Mettre bord haut
	opt.GeoM.Translate(-float64(width)+23+3.4, -20)
	opt.GeoM.Scale(6.9, 1)
	imgOut.DrawImage(imgTab[1], opt)
	return imgOut
}
