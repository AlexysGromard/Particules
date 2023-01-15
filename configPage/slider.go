package configPage

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Creation de la structure du slider
type Slider struct {
	x, y, width, height int
	sliderXPosition     int
	imageSliderTrack    *ebiten.Image
	imageSlider         *ebiten.Image
	imageSliderHover    *ebiten.Image
	imageSliderDisabled *ebiten.Image
	hover               bool
	pressed             bool
	value               *float64
	minValue            float64
	maxValue            float64
	disabled            bool
}

// Impression des sliders
// Cette fonction met à jour l'affichage des sliders en fonction de leur état (hover, pressed, checked)
func (s *Slider) Draw(screen *ebiten.Image) {
	var img *ebiten.Image
	if s.disabled {
		img = s.imageSliderDisabled
	} else if s.hover {
		img = s.imageSliderHover
	} else {
		img = s.imageSlider
	}
	// Draw imageSliderTrack
	optSliderTrack := &ebiten.DrawImageOptions{}
	scaleX := float64(s.width) / float64(s.imageSliderTrack.Bounds().Dx())
	optSliderTrack.GeoM.Scale(scaleX, 1)
	optSliderTrack.GeoM.Translate(float64(s.x), float64(s.y))
	// Draw imageSlider
	screen.DrawImage(s.imageSliderTrack, optSliderTrack)
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(float64(4), 1)
	opt.GeoM.Translate(float64(s.sliderXPosition), float64(s.y))
	screen.DrawImage(img, opt)
	// Si le slider est hover, alors cursor = pointer
	if s.hover {
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
	} else {
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	}
}

// Mise à jour des sliders
// Cette fonction regarde si la souris est sur un slider et si elle est enfoncée
// Si le slider est enfoncé, il est déplacé
func (s *Slider) Update(screen *ebiten.Image) {
	// Position de la souris
	x, y := ebiten.CursorPosition()
	// Position de la souris par rapport au slider
	difX := x - s.sliderXPosition
	s.hover = false
	// Si la souris est sur le slider
	if ((x >= s.sliderXPosition-difX && x <= s.sliderXPosition+5*4 && y >= s.y && y <= s.y+s.height+4) || s.pressed) && !s.disabled {
		s.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			s.pressed = true
			// Bouger en fonction du mouvement de la souris
			if x > s.sliderXPosition {
				s.sliderXPosition -= s.x + (s.sliderXPosition - x)
			} else {
				s.sliderXPosition += s.x + (x - s.sliderXPosition)
			}
			// Limiter le slider à la taille de la barre
			if s.sliderXPosition < s.x {
				s.sliderXPosition = s.x
			} else if s.sliderXPosition+5*4 > s.x+s.width {
				s.sliderXPosition = s.x + s.width - 5*4
			}
		} else {
			s.pressed = false
		}
	}
	// Mettre à jour la valeur du slider au pointeur
	*s.value = float64(s.sliderXPosition-s.x) / float64(s.width) * s.maxValue
}

// Création d'un slider
// Cette fonction crée un slider
func newSlider(x, y, width, height int, imageSliderTrack, imageSlider, imageSliderHover, imageSliderDisabled *ebiten.Image, value *float64, minValue float64, maxValue float64, disabled bool) *Slider {
	return &Slider{
		x:                   x,
		y:                   y,
		width:               width,
		height:              height,
		sliderXPosition:     x + int(float64(width)*(*value/maxValue)),
		imageSliderTrack:    imageSliderTrack,
		imageSlider:         imageSlider,
		imageSliderHover:    imageSliderHover,
		imageSliderDisabled: imageSliderDisabled,
		value:               value,
		minValue:            minValue,
		maxValue:            maxValue,
		disabled:            disabled,
	}
}
