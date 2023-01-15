package configPage

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Creation de la structure du slider
type SliderF struct {
	x, y, width, height int
	sliderXPosition     float64
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
	round               bool
}

type SliderI struct {
	x, y, width, height int
	sliderXPosition     float64
	imageSliderTrack    *ebiten.Image
	imageSlider         *ebiten.Image
	imageSliderHover    *ebiten.Image
	imageSliderDisabled *ebiten.Image
	hover               bool
	pressed             bool
	value               *int
	minValue            float64
	maxValue            float64
	disabled            bool
}

// Impression des sliders
// Cette fonction met à jour l'affichage des sliders en fonction de leur état (hover, pressed, checked)
func (s *SliderF) Draw(screen *ebiten.Image) {
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
}
func (s *SliderI) Draw(screen *ebiten.Image) {
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
}

// Mise à jour des sliders
// Cette fonction regarde si la souris est sur un slider et si elle est enfoncée
// Si le slider est enfoncé, il est déplacé
func (s *SliderF) Update(screen *ebiten.Image) {
	// Position de la souris
	x, y := ebiten.CursorPosition()
	// Position de la souris par rapport au slider
	difX := float64(x) - s.sliderXPosition
	s.hover = false
	// Si la souris est sur le slider
	if ((float64(x) >= s.sliderXPosition-difX && float64(x) <= s.sliderXPosition+5*4 && y >= s.y && y <= s.y+s.height+4) || s.pressed) && !s.disabled {
		s.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			s.pressed = true
			// Bouger en fonction du mouvement de la souris
			if float64(x) > s.sliderXPosition {
				s.sliderXPosition -= (s.sliderXPosition - float64(x))
			} else {
				s.sliderXPosition += (float64(x) - s.sliderXPosition)
			}
			// Limiter le slider à la taille de la barre
			if s.sliderXPosition < float64(s.x) {
				s.sliderXPosition = float64(s.x)
			} else if s.sliderXPosition+5*4 > float64(s.x+s.width) {
				s.sliderXPosition = float64(s.x + s.width - 5*4)
			}
		} else {
			s.pressed = false
		}
	}
	// Mettre à jour la valeur du slider au pointeur
	// *s.value = (float64(s.sliderXPosition-s.x)/float64(s.width))*s.maxValue + s.minValue
	*s.value = (float64(s.sliderXPosition-float64(s.x))/float64(s.width-20))*(s.maxValue-s.minValue) + s.minValue

	// Arrondir la valeur du slider
	if s.round {
		*s.value = float64(int(*s.value))
	}
}
func (s *SliderI) Update(screen *ebiten.Image) {
	// Position de la souris
	x, y := ebiten.CursorPosition()
	// Position de la souris par rapport au slider
	difX := float64(x) - s.sliderXPosition
	s.hover = false
	// Si la souris est sur le slider
	if ((float64(x) >= s.sliderXPosition-difX && float64(x) <= s.sliderXPosition+5*4 && y >= s.y && y <= s.y+s.height+4) || s.pressed) && !s.disabled {
		s.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			s.pressed = true
			// Bouger en fonction du mouvement de la souris
			if float64(x) > s.sliderXPosition {
				s.sliderXPosition -= (s.sliderXPosition - float64(x))
			} else {
				s.sliderXPosition += (float64(x) - s.sliderXPosition)
			}
			// Limiter le slider à la taille de la barre
			if int(s.sliderXPosition) < s.x {
				s.sliderXPosition = float64(s.x)
			} else if s.sliderXPosition+5*4 > float64(s.x+s.width) {
				s.sliderXPosition = float64(s.x + s.width - 5*4)
			}
		} else {
			s.pressed = false
		}
	}
	// Mettre à jour la valeur du slider au pointeur
	*s.value = int((float64(s.sliderXPosition-float64(s.x))/float64(s.width-20))*(s.maxValue-s.minValue) + s.minValue)
}

// Création d'un slider
// Cette fonction crée un slider
func newSliderFloat(x, y, width, height int, images []*ebiten.Image, value *float64, minValue float64, maxValue float64, disabled bool, round bool) *SliderF {
	return &SliderF{
		x:                   x,
		y:                   y,
		width:               width,
		height:              height,
		sliderXPosition:     (float64(*value-minValue)/(maxValue-minValue))*float64(width-20) + float64(x),
		imageSliderTrack:    images[0],
		imageSlider:         images[1],
		imageSliderHover:    images[2],
		imageSliderDisabled: images[3],
		value:               value,
		minValue:            minValue,
		maxValue:            maxValue,
		disabled:            disabled,
		round:               round,
	}
}

func newSliderInt(x, y, width, height int, images []*ebiten.Image, value *int, minValue float64, maxValue float64, disabled bool) *SliderI {
	return &SliderI{
		x:                   x,
		y:                   y,
		width:               width,
		height:              height,
		sliderXPosition:     (float64(float64(*value)-minValue)/(maxValue-minValue))*float64(width-20) + float64(x),
		imageSliderTrack:    images[0],
		imageSlider:         images[1],
		imageSliderHover:    images[2],
		imageSliderDisabled: images[3],
		value:               value,
		minValue:            minValue,
		maxValue:            maxValue,
		disabled:            disabled,
	}
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}
