package configPage

import (
	"image/color"

	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

// NOTICE DES COMMPOSANTS
// Text : positionX, positionY, text, PoliceDécriture, couleur(RGBA)
// Button : positionX, positionY, largeur, hauteur, images []*ebiten.Image, text, PoliceDécriture, fonction à appeler
// ----- images []*ebiten.Image = [0]normal, [1]pressed, [2]hover
// Checkbox : positionX, positionY, largeur, hauteur, images []*ebiten.Image, valeur(checked), disabled, fonction à appeler
// ----- images []*ebiten.Image = [0]normal, [1]hover, [2]disabled, [3]checked, [4]unchecked, [5]greyed(disabled)
// Slider : positionX, positionY, largeur, hauteur, imageLigneSlider, imageDuSlider, imageDuSliderHover, imageDuSliderDisabled, valeur(pointeur vers var), valeurMin, valeurMax, disabled
// numberInput : positionX, positionY, largeur, hauteur, imageNormal, imageHover, imageDisabled, disabled, number(pointeur), PoliceDécriture

var (
	welcomeTitle          *Text
	accessParticlesButton *Button
	textInut              *NumberInput
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
		accessParticlesButton = newButton(config.General.WindowSizeX-150-30, config.General.WindowSizeY-50-30, 150, 50, images, "Valider", RobotoRegularFontF, func() { SaveConfig() })
	}
	// Met à jour l'état du bouton et l'affiche
	accessParticlesButton.Update(screen)
	accessParticlesButton.Draw(screen)

	// Text input
	if textInut == nil {
		textInut = newTextInput(10, 300, 100, 50, findImage(ImageList, "text-input-idle.png"), findImage(ImageList, "text-input-hover.png"), findImage(ImageList, "text-input-disabled.png"), true, &config.General.SpawnX, RobotoRegularFontF)
	}
	// Update text input
	textInut.Update(screen)
	textInut.Draw(screen)

	// Crée la checkbox si elle n'existe pas
	if checkbox == nil {
		images := []*ebiten.Image{findImage(ImageList, "checkbox-idle.png"), findImage(ImageList, "checkbox-hover.png"), findImage(ImageList, "checkbox-disabled.png"), findImage(ImageList, "checkbox-checked-idle.png"), findImage(ImageList, "checkbox-unchecked-idle.png"), findImage(ImageList, "checkbox-greyed-disabled.png")}
		checkbox = newCheckbox(10, 70, 50, 50, images, config.General.Debug, true, func() { config.General.Debug = !config.General.Debug })
	}
	// Met à jour l'état de la checkbox et l'affiche
	checkbox.Update(screen)
	checkbox.Draw(screen)

	// Create text
	if texttest == nil {
		texttest = newText(70, 90, "Debug", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	// Draw text
	texttest.Draw(screen)

	// Create slider
	if sliderTest == nil {
		sliderTest = newSlider(10, 150, 200, 3, findImage(ImageList, "slider-track-idle.png"), findImage(ImageList, "slider-handle-idle.png"), findImage(ImageList, "slider-handle-hover.png"), findImage(ImageList, "slider-handle-disabled.png"), &config.General.Gravity, 0, 1, true)
	}
	// Update slider
	sliderTest.Update(screen)
	// Draw slider
	sliderTest.Draw(screen)

	return nil
}
