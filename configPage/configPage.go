package configPage

import (
	"fmt"

	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

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
