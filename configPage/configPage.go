package configPage

import (
	"fmt"

	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
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

// Création de la structure Checkbox
type Checkbox struct {
	x, y, width, height int
	imageNormal         *ebiten.Image
	imageHover          *ebiten.Image
	circleTrue          *ebiten.Image
	circleFalse         *ebiten.Image
	checked             bool
	hover               bool
	pressed             bool
	justPressed         bool
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
}

// Impression des Checkbox
// Cette fonction met à jour l'affichage des checkbox en fonction de leur état (hover, pressed, checked)
func (c *Checkbox) Draw(screen *ebiten.Image) {
	var img *ebiten.Image
	if c.hover {
		img = c.imageHover
	} else {
		img = c.imageNormal
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(img, opt)
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
		b.x = config.General.WindowSizeX - 100 - 10
		b.y = config.General.WindowSizeY - 50 - 10
	}
}

// Mise à jour des checkbox
// Cette fonction regarde si la souris est sur une checkbox et si elle est enfoncée
// Elle regarde vérifie si la checkbox est cochée ou non
// Si la checkbox est cochée, elle affiche l'image
func (c *Checkbox) updateCheckbox(screen *ebiten.Image) {
	// Position de la souris
	x, y := ebiten.CursorPosition()
	c.hover = false
	// Si la souris est sur la checkbox
	if x > c.x && x < c.x+c.width && y > c.y && y < c.y+c.height {
		c.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if !c.justPressed {
				c.checked = !c.checked
				c.justPressed = true
				c.onClick()
			}
		} else {
			c.pressed = false
			c.justPressed = false

		}
	}
	// Si la checkbox est cochée
	if c.checked {
		// Afficher l'image de droite
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(c.x), float64(c.y))
		screen.DrawImage(c.circleTrue, opt)
	} else {
		// Afficher l'image de gauche
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(c.x), float64(c.y))
		screen.DrawImage(c.circleFalse, opt)
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
		textX:        x + width/2,
		textY:        y + height/2,
		font:         font,
		onClick:      onClick,
	}
}

// NewCheckbox crée une nouvelle checkbox en fonction des paramètres
// Entrées : x,y (position), width, height (taille), images (liste des images), checked (état de la checkbox), onClick (fonction à appeler quand on clique sur la checkbox)
func newCheckbox(x, y, width, height int, images []*ebiten.Image, checked bool, onClick func()) *Checkbox {
	return &Checkbox{
		x:           x,
		y:           y,
		width:       width,
		height:      height,
		imageNormal: images[0],
		imageHover:  images[1],
		circleTrue:  images[2],
		circleFalse: images[3],
		checked:     checked,
		onClick:     onClick,
	}
}

var (
	accessParticlesButton *Button
	checkbox              *Checkbox
)

// FidnImage cherche une image grâce à son nom dans la liste des images
func findImage(images []Images, filename string) *ebiten.Image {
	// Boucle dans les images avec leurs noms
	for _, img := range images {
		// Trouve l'image avec le nom correspondant
		if img.name == filename {
			// Return img
			return img.image
		}
	}
	return nil
}

// UpdateConfigPage est la fonction qui est appelée pour mettre à jour la page de configuration
func UpdateConfigPage(screen *ebiten.Image) error {
	fontFace, _ := loadFont(fontFaceRegular, 20)
	// Crée le boutton si il n'existe pas
	if accessParticlesButton == nil {
		images := []*ebiten.Image{findImage(ImageList, "button-idle.png"), findImage(ImageList, "button-pressed.png"), findImage(ImageList, "button-hover.png")}
		accessParticlesButton = newButton(config.General.WindowSizeX-100-10, config.General.WindowSizeY-50-10, 100, 50, images, "Hello World", fontFace, func() { fmt.Println("test") })
	}
	// Met à jour l'état du bouton et l'affiche
	accessParticlesButton.updateButton(screen)
	accessParticlesButton.Draw(screen)

	// Crée la checkbox si elle n'existe pas
	if checkbox == nil {
		images := []*ebiten.Image{findImage(ImageList, "checkbox-idle.png"), findImage(ImageList, "checkbox-hover.png"), findImage(ImageList, "checkbox-checked-idle.png"), findImage(ImageList, "checkbox-unchecked-idle.png")}
		checkbox = newCheckbox(10, 70, 50, 50, images, config.General.Debug, func() { config.General.Debug = !config.General.Debug })
	}
	// Met à jour l'état de la checkbox et l'affiche
	checkbox.updateCheckbox(screen)
	checkbox.Draw(screen)

	return nil
}
