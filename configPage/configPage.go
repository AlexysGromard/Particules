package configPage

import (
	"fmt"
	"image/color"

	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	welcomeTitle          *Text
	accessParticlesButton *Button
	textInut              *TextInput
	checkbox              *Checkbox
	texttest              *Text
	sliderTest            *Slider
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
	// Load les fonts
	err := loadFontRegular()
	if err != nil {
		return err
	}
	err = loadFontBold()
	if err != nil {
		return err
	}
	err = loadFontTitle()
	if err != nil {
		return err
	}
	// Titre de configuration
	if welcomeTitle == nil {
		welcomeTitle = newText(10, 30, "Configuration", PageTitleF, color.RGBA{127, 139, 148, 255})
	}
	welcomeTitle.Draw(screen)

	// Crée le boutton si il n'existe pas
	if accessParticlesButton == nil {
		images := []*ebiten.Image{findImage(ImageList, "button-idle.png"), findImage(ImageList, "button-pressed.png"), findImage(ImageList, "button-hover.png")}
		accessParticlesButton = newButton(config.General.WindowSizeX-120-10, config.General.WindowSizeY-50-10, 120, 50, images, "Hello World", RobotoRegularFontF, func() { fmt.Println("test") })
	}
	// Met à jour l'état du bouton et l'affiche
	accessParticlesButton.updateButton(screen)
	accessParticlesButton.Draw(screen)

	// Text input
	if textInut == nil {
		textInut = newTextInput(10, 300, 200, 50, findImage(ImageList, "text-input-idle.png"), findImage(ImageList, "text-input-hover.png"), findImage(ImageList, "tool-tip.png"), &config.General.SpawnX, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	// Update text input
	textInut.updateTextInput(screen)
	textInut.Draw(screen)

	// Crée la checkbox si elle n'existe pas
	if checkbox == nil {
		images := []*ebiten.Image{findImage(ImageList, "checkbox-idle.png"), findImage(ImageList, "checkbox-hover.png"), findImage(ImageList, "checkbox-checked-idle.png"), findImage(ImageList, "checkbox-unchecked-idle.png")}
		checkbox = newCheckbox(10, 70, 50, 50, images, config.General.Debug, func() { config.General.Debug = !config.General.Debug })
	}
	// Met à jour l'état de la checkbox et l'affiche
	checkbox.updateCheckbox(screen)
	checkbox.Draw(screen)

	// Create text
	if texttest == nil {
		texttest = newText(70, 90, "Debug", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	// Draw text
	texttest.Draw(screen)

	// Create slider
	if sliderTest == nil {
		sliderTest = newSlider(10, 150, 200, 3, findImage(ImageList, "slider-track-idle.png"), findImage(ImageList, "slider-handle-idle.png"), findImage(ImageList, "slider-handle-hover.png"), findImage(ImageList, "slider-handle-disabled.png"), &config.General.Gravity, 0, 1)
	}
	// Update slider
	sliderTest.updateSlider(screen)
	// Draw slider
	sliderTest.Draw(screen)

	return nil
}
