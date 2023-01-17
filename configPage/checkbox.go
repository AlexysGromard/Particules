package configPage

import "github.com/hajimehoshi/ebiten/v2"

// Création de la structure Checkbox
type Checkbox struct {
	x, y, width, height int
	imageNormal         *ebiten.Image
	imageHover          *ebiten.Image
	imageDisabled       *ebiten.Image
	circleTrue          *ebiten.Image
	circleFalse         *ebiten.Image
	circleDisabled      *ebiten.Image
	checked             bool
	hover               bool
	pressed             bool
	justPressed         bool
	disabled            bool
	onClick             func()
}

// Impression des Checkbox
// Cette fonction met à jour l'affichage des checkbox en fonction de leur état (hover, pressed, checked)
func (c *Checkbox) draw(screen *ebiten.Image) {
	var img *ebiten.Image
	if c.hover && !c.disabled {
		img = c.imageHover
	} else if c.disabled {
		img = c.imageDisabled
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
func (c *Checkbox) update(screen *ebiten.Image) {
	// Position de la souris
	x, y := ebiten.CursorPosition()
	c.hover = false
	// Si la souris est sur la checkbox
	if x > c.x && x < c.x+c.width && y > c.y && y < c.y+c.height && !c.disabled {
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
	if c.disabled {
		// Afficher l'image disabled
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(c.x), float64(c.y))
		screen.DrawImage(c.circleDisabled, opt)
	} else if c.checked {
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
	// Ajouter le scroll
	c.y += ScrollY
	// Draw la checkbox
	c.draw(screen)
}

// NewCheckbox crée une nouvelle checkbox en fonction des paramètres
// Entrées : x,y (position), width, height (taille), images (liste des images), checked (état de la checkbox), onClick (fonction à appeler quand on clique sur la checkbox)
func newCheckbox(x, y, width, height int, images []*ebiten.Image, checked bool, disabled bool, onClick func()) *Checkbox {
	return &Checkbox{
		x:              x,
		y:              y,
		width:          width,
		height:         height,
		imageNormal:    images[0],
		imageHover:     images[1],
		imageDisabled:  images[2],
		circleTrue:     images[3],
		circleFalse:    images[4],
		circleDisabled: images[5],
		checked:        checked,
		disabled:       disabled,
		onClick:        onClick,
	}
}
