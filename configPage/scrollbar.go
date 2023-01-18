package configPage

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

// Creation de la structure de la scrollbar
type Scrollbar struct {
	// Scrollbar position
	x, y float64
	// Scrollbar size
	width, height int
	// Images
	imageNormal *ebiten.Image
	imageHover  *ebiten.Image
	// Scrollbar state
	pressed    bool
	hover      bool
	crollValue float64
}

// Impression de la scrollbar
// Cette fonction met à jour l'affichage de la scrollbar en fonction de son état (hover, pressed)
func (s *Scrollbar) draw(screen *ebiten.Image) {
	var img *ebiten.Image
	if s.hover {
		img = s.imageHover
	} else {
		img = s.imageNormal
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(s.x), float64(s.y))
	opt.GeoM.Scale(1, float64(s.height)/5)
	screen.DrawImage(nineSlice(img, s.width, s.height), opt)
}

// Mise à jour de la scrollbar
// Cette fonction met à jour l'état de la scrollbar en fonction de la position de la souris
// Elle change aussi sa position en fonction du scroll
func (s *Scrollbar) update(screen *ebiten.Image) {
	// Position de la souris
	x, y := ebiten.CursorPosition()
	s.hover = false
	// Si la souris est sur le boutton
	if x > int(s.x) && x < int(s.x)+s.width && y > int(s.y) && y < int(s.y)+s.height {
		s.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			s.pressed = true
		} else {
			s.pressed = false
		}
	} else {
		s.pressed = false
	}
	// Variable pour mettre à jour la valeur de défilement d'un élément
	s.crollValue += -float64(ScrollY) / float64(config.General.WindowSizeY)
	// Modifie la position de la scrollbar (y)
	s.y = float64(s.crollValue) / float64(s.width/5) * float64(s.width)
	s.draw(screen)
}

// Création d'une scrollbar
func newScrollbar(x, y float64, width, height int, images []*ebiten.Image) *Scrollbar {
	return &Scrollbar{
		x:           x,
		y:           y,
		width:       width,
		height:      height,
		imageNormal: images[0],
		imageHover:  images[1],
	}
}
