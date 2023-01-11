package configPage

import (
	"fmt"

	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

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

func (b *Button) updateButton(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	b.hover = false
	if x > b.x && x < b.x+b.width && y > b.y && y < b.y+b.height {
		b.hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			b.pressed = true
		}
	} else {
		b.pressed = false
	}
	if b.x < config.General.WindowSizeX-b.width || b.x > config.General.WindowSizeX-b.width+10 || b.y < config.General.WindowSizeY-b.height || b.y > config.General.WindowSizeY-b.height+10 {
		b.x = config.General.WindowSizeX - 100 - 10
		b.y = config.General.WindowSizeY - 50 - 10
	}
}

func (c *Checkbox) updateCheckbox(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	c.hover = false
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
	if c.checked {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(c.x), float64(c.y))
		screen.DrawImage(c.circleTrue, opt)
	} else {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(c.x), float64(c.y))
		screen.DrawImage(c.circleFalse, opt)
	}

}

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
	button   *Button
	checkbox *Checkbox
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

func UpdateConfigPage(screen *ebiten.Image) error {
	fontFace, _ := loadFont(fontFaceRegular, 20)
	// Initialize button if it hasn't been created yet
	if button == nil {
		images := []*ebiten.Image{findImage(ImageList, "button-idle.png"), findImage(ImageList, "button-pressed.png"), findImage(ImageList, "button-hover.png")}
		button = newButton(config.General.WindowSizeX-100-10, config.General.WindowSizeY-50-10, 100, 50, images, "Hello World", fontFace, func() { fmt.Println("test") })
	}
	// Update the button's state
	button.updateButton(screen)
	// Draw the button
	button.Draw(screen)
	// Create checkbox if it hasn't been created yet
	if checkbox == nil {
		images := []*ebiten.Image{findImage(ImageList, "checkbox-idle.png"), findImage(ImageList, "checkbox-hover.png"), findImage(ImageList, "checkbox-checked-idle.png"), findImage(ImageList, "checkbox-unchecked-idle.png")}
		checkbox = newCheckbox(10, 70, 50, 50, images, config.General.Debug, func() { config.General.Debug = !config.General.Debug })
	}
	checkbox.updateCheckbox(screen)
	checkbox.Draw(screen)

	return nil
}
