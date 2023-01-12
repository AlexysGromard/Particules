package configPage

import "github.com/hajimehoshi/ebiten/v2"

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
