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
// Checkbox : positionX, positionY, largeur, hauteur, images []*ebiten.Image, valeur(checked), disabled(pointeur), fonction à appeler
// ----- images []*ebiten.Image = [0]normal, [1]hover, [2]disabled, [3]checked, [4]unchecked, [5]greyed(disabled)
// Slider : positionX, positionY, largeur, hauteur, imageLigneSlider, imageDuSlider, imageDuSliderHover, imageDuSliderDisabled, valeur(pointeur vers var), valeurMin, valeurMax, disabled
// numberInput : positionX, positionY, largeur, hauteur, imageNormal, imageHover, imageDisabled, disabled, number(pointeur), PoliceDécriture

// POLICES DE CARACTÈRES
// RobotoRegularFontF : Roboto Regular 20
// RobotoBoldFontF : Roboto Bold 20
// PageTitleF : Roboto Bold 25

var (
	welcomeTitle *Text
	// windowConfiguration
	windowConfiguration     *Text
	debugButton             *Checkbox
	DebugText               *Text
	particleGeneration      *Text
	initNumParticles        *NumberInput
	initNumParticlesText    *Text
	randomSpawn             *Checkbox
	randomSpawnText         *Text
	spawnCenter             *Checkbox
	spawnCenterText         *Text
	spawnX                  *NumberInput
	spawnXText              *Text
	spawnY                  *NumberInput
	spawnYText              *Text
	spawnRate               *Slider
	spawnRateText           *Text
	spawnRateValue          *Value
	spawnOnAnObject         *Checkbox
	spawnOnAnObjectText     *Text
	spawnOnObjectCircle     *Checkbox
	spawnOnObjectCircleText *Text
	spawnOnObjectSquare     *Checkbox
	spawnOnObjectSquareText *Text
	spawnObjectWidth        *NumberInput
	spawnObjectWidthText    *Text
	//particulesPropertiesText
	particulesPropertiesText *Text
	rotation                 *Slider
	rotationValue            *Value
	rotationText             *Text
	scaleX                   *Slider
	scaleXValue              *Value
	scaleXText               *Text
	scaleY                   *Slider
	scaleYValue              *Value
	scaleYText               *Text
	opacity                  *Slider
	opacityValue             *Value
	opacityText              *Text
	colorRed                 *Slider
	colorRedValue            *Value
	colorRedText             *Text
	colorGreen               *Slider
	colorGreenValue          *Value
	colorGreenText           *Text

	accessParticlesButton *Button
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
	// EbitenImages
	var (
		// Checkbox
		checkboxImages = []*ebiten.Image{
			findImage(ImageList, "checkbox-idle.png"),            // Default
			findImage(ImageList, "checkbox-hover.png"),           // Hover
			findImage(ImageList, "checkbox-disabled.png"),        // Disabled
			findImage(ImageList, "checkbox-checked-idle.png"),    // Checked
			findImage(ImageList, "checkbox-unchecked-idle.png"),  // Unchecked
			findImage(ImageList, "checkbox-greyed-disabled.png")} // Greyed(Disabled)
		// Button
		buttonImages = []*ebiten.Image{
			findImage(ImageList, "button-idle.png"),    // Default
			findImage(ImageList, "button-pressed.png"), // Pressed
			findImage(ImageList, "button-hover.png")}   // Hover
		// Slider
		sliderImages = []*ebiten.Image{
			findImage(ImageList, "slider-track-idle.png"),      // Line
			findImage(ImageList, "slider-handle-idle.png"),     // Default
			findImage(ImageList, "slider-handle-hover.png"),    // Hover
			findImage(ImageList, "slider-handle-disabled.png")} // Disabled
		// NumberInput
		numberInputImages = []*ebiten.Image{
			findImage(ImageList, "text-input-idle.png"),     // Default
			findImage(ImageList, "text-input-hover.png"),    // Hover
			findImage(ImageList, "text-input-disabled.png")} // Disabled
	)
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
		welcomeTitle = newText(20, 30, "Configuration", PageTitleF, color.RGBA{127, 139, 148, 255})
	}
	welcomeTitle.Draw(screen)

	// Titre de configuration de la fenêtre
	if windowConfiguration == nil {
		windowConfiguration = newText(20, 60, "Configuration de la fenêtre", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	}
	windowConfiguration.Draw(screen)

	// Checkbox debug
	if debugButton == nil {
		debugButton = newCheckbox(20, 75, 50, 30, checkboxImages, config.General.Debug, false, func() { config.General.Debug = !config.General.Debug })
	}
	debugButton.Update(screen)
	debugButton.Draw(screen)
	if DebugText == nil {
		DebugText = newText(100, 95, "Activer le mode debug", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	DebugText.Draw(screen)

	// Titre de génération de particules
	if particleGeneration == nil {
		particleGeneration = newText(20, 150, "Génération de particules", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	}
	particleGeneration.Draw(screen)

	// Nombre de particules initiales
	if initNumParticles == nil {
		initNumParticles = newTextInput(20, 165, 100, 30, numberInputImages, false, &config.General.InitNumParticles, RobotoRegularFontF)
	}
	initNumParticles.Update(screen)
	initNumParticles.Draw(screen)
	if initNumParticlesText == nil {
		initNumParticlesText = newText(110, 185, "Nombre de particules initiales", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	initNumParticlesText.Draw(screen)

	// Checkbox random spawn
	if randomSpawn == nil {
		randomSpawn = newCheckbox(20, 200, 50, 30, checkboxImages, config.General.RandomSpawn, false, func() { config.General.RandomSpawn = !config.General.RandomSpawn })
	}
	randomSpawn.Update(screen)
	randomSpawn.Draw(screen)
	if randomSpawnText == nil {
		randomSpawnText = newText(110, 220, "Activer le spawn aléatoire", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	randomSpawnText.Draw(screen)

	// Checkbox spawn center
	if spawnCenter == nil {
		spawnCenter = newCheckbox(20, 235, 50, 30, checkboxImages, config.General.SpawnCenter, config.General.RandomSpawn, func() { config.General.SpawnCenter = !config.General.SpawnCenter })
	}
	spawnCenter.disabled = config.General.RandomSpawn
	spawnCenter.Update(screen)
	spawnCenter.Draw(screen)
	if spawnCenterText == nil {
		spawnCenterText = newText(110, 255, "Activer le spawn au centre", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnCenterText.Draw(screen)

	// Input de la position de spawnX
	if spawnX == nil {
		spawnX = newTextInput(20, 270, 100, 30, numberInputImages, config.General.SpawnCenter, &config.General.SpawnX, RobotoRegularFontF)
	}
	spawnX.disabled = config.General.SpawnCenter
	spawnX.Update(screen)
	spawnX.Draw(screen)
	if spawnXText == nil {
		spawnXText = newText(110, 290, "Position de spawn X", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnXText.Draw(screen)

	// Input de la position de spawnY
	if spawnY == nil {
		spawnY = newTextInput(20, 305, 100, 30, numberInputImages, config.General.SpawnCenter, &config.General.SpawnY, RobotoRegularFontF)
	}
	spawnY.disabled = config.General.SpawnCenter
	spawnY.Update(screen)
	spawnY.Draw(screen)
	if spawnYText == nil {
		spawnYText = newText(110, 325, "Position de spawn Y", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnYText.Draw(screen)

	// Input SpawnRate Slider
	if spawnRate == nil {
		spawnRate = newSlider(20, 350, 100, 5, sliderImages, &config.General.SpawnRate, 0, 100, false, true)
	}
	spawnRate.Update(screen)
	spawnRate.Draw(screen)
	if spawnRateText == nil {
		spawnRateText = newText(165, 360, "Fréquence de génération", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnRateText.Draw(screen)
	if spawnRateValue == nil {
		spawnRateValue = newValue(125, 360, &config.General.SpawnRate, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnRateValue.Update(screen)
	spawnRateValue.Draw(screen)

	// Checkbox SpawnOnAnObject
	if spawnOnAnObject == nil {
		spawnOnAnObject = newCheckbox(20, 375, 50, 30, checkboxImages, config.General.SpawnOnAnObject, false, func() { config.General.SpawnOnAnObject = !config.General.SpawnOnAnObject })
	}
	spawnOnAnObject.Update(screen)
	spawnOnAnObject.Draw(screen)
	if spawnOnAnObjectText == nil {
		spawnOnAnObjectText = newText(110, 395, "Activer le spawn sur un objet", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnOnAnObjectText.Draw(screen)

	// Checkbox SpawnObject = "circle"
	if spawnOnObjectCircle == nil {
		spawnOnObjectCircle = newCheckbox(40, 410, 50, 30, checkboxImages, config.General.SpawnObject == "circle", !config.General.SpawnOnAnObject, func() { config.General.SpawnObject = "circle" })
	}
	spawnOnObjectCircle.checked = config.General.SpawnObject == "circle"
	spawnOnObjectCircle.Update(screen)
	spawnOnObjectCircle.Draw(screen)
	if spawnOnObjectCircleText == nil {
		spawnOnObjectCircleText = newText(110, 430, "Cercle", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnOnObjectCircleText.Draw(screen)

	// Checkbox SpawnObject = "square"
	if spawnOnObjectSquare == nil {
		spawnOnObjectSquare = newCheckbox(40, 445, 50, 30, checkboxImages, config.General.SpawnObject == "square", !config.General.SpawnOnAnObject, func() { config.General.SpawnObject = "square" })
	}
	spawnOnObjectSquare.checked = config.General.SpawnObject == "square"
	spawnOnObjectSquare.Update(screen)
	spawnOnObjectSquare.Draw(screen)
	if spawnOnObjectSquareText == nil {
		spawnOnObjectSquareText = newText(110, 465, "Carré", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnOnObjectSquareText.Draw(screen)
	// (Vérifie que SpawnOnAnObject est activé)
	spawnOnObjectCircle.disabled, spawnOnObjectSquare.disabled = !config.General.SpawnOnAnObject, !config.General.SpawnOnAnObject

	// Input spawnObjectWidth number
	if spawnObjectWidth == nil {
		spawnObjectWidth = newTextInput(20, 480, 100, 30, numberInputImages, !config.General.SpawnOnAnObject, &config.General.SpawnObjectWidth, RobotoRegularFontF)
	}
	spawnObjectWidth.disabled = !config.General.SpawnOnAnObject
	spawnObjectWidth.Update(screen)
	spawnObjectWidth.Draw(screen)
	if spawnObjectWidthText == nil {
		spawnObjectWidthText = newText(110, 500, "Largeur de l'objet", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	spawnObjectWidthText.Draw(screen)

	// Titre particulesPropertiesText
	if particulesPropertiesText == nil {
		particulesPropertiesText = newText(20, 550, "Propriétés des particules", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	}
	particulesPropertiesText.Draw(screen)

	// Slider rotation
	if rotation == nil {
		rotation = newSlider(20, 575, 100, 5, sliderImages, &config.General.Rotation, 0, 3.14159*2, false, false)
	}
	rotation.Update(screen)
	rotation.Draw(screen)
	if rotationValue == nil {
		rotationValue = newValue(125, 585, &config.General.Rotation, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	rotationValue.Update(screen)
	rotationValue.Draw(screen)

	if rotationText == nil {
		rotationText = newText(175, 585, "Rotation", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	rotationText.Draw(screen)

	// Slider ScaleX
	if scaleX == nil {
		scaleX = newSlider(20, 600, 100, 5, sliderImages, &config.General.ScaleX, 0, 5, false, true)
	}
	scaleX.Update(screen)
	scaleX.Draw(screen)
	if scaleXValue == nil {
		scaleXValue = newValue(125, 610, &config.General.ScaleX, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	scaleXValue.Update(screen)
	scaleXValue.Draw(screen)
	if scaleXText == nil {
		scaleXText = newText(175, 610, "ScaleX", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	scaleXText.Draw(screen)
	// Slider ScaleY
	if scaleY == nil {
		scaleY = newSlider(20, 625, 100, 5, sliderImages, &config.General.ScaleY, 0, 5, false, true)
	}
	scaleY.Update(screen)
	scaleY.Draw(screen)
	if scaleYValue == nil {
		scaleYValue = newValue(125, 635, &config.General.ScaleY, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	scaleYValue.Update(screen)
	scaleYValue.Draw(screen)
	if scaleYText == nil {
		scaleYText = newText(175, 635, "ScaleY", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	scaleYText.Draw(screen)

	// Opacity
	if opacity == nil {
		opacity = newSlider(20, 650, 100, 5, sliderImages, &config.General.Opacity, 0, 1, false, false)
	}
	opacity.Update(screen)
	opacity.Draw(screen)
	if opacityValue == nil {
		opacityValue = newValue(125, 660, &config.General.Opacity, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	opacityValue.Update(screen)
	opacityValue.Draw(screen)
	if opacityText == nil {
		opacityText = newText(175, 660, "Opacité", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	opacityText.Draw(screen)

	// ColorRed
	if colorRed == nil {
		colorRed = newSlider(20, 675, 100, 5, sliderImages, &config.General.ColorRed, 0, 1, false, false)
	}
	colorRed.Update(screen)
	colorRed.Draw(screen)
	if colorRedValue == nil {
		colorRedValue = newValue(125, 685, &config.General.ColorRed, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	colorRedValue.Update(screen)
	colorRedValue.Draw(screen)
	if colorRedText == nil {
		colorRedText = newText(175, 685, "Rouge", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	colorRedText.Draw(screen)

	// ColorGreen
	if colorGreen == nil {
		colorGreen = newSlider(20, 700, 100, 5, sliderImages, &config.General.ColorGreen, 0, 1, false, false)
	}
	colorGreen.Update(screen)
	colorGreen.Draw(screen)
	if colorGreenValue == nil {
		colorGreenValue = newValue(125, 710, &config.General.ColorGreen, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	colorGreenValue.Update(screen)
	colorGreenValue.Draw(screen)
	if colorGreenText == nil {
		colorGreenText = newText(175, 710, "Vert", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	}
	colorGreenText.Draw(screen)

	// Crée le boutton si il n'existe pas
	if accessParticlesButton == nil {
		accessParticlesButton = newButton(config.General.WindowSizeX-150-30, config.General.WindowSizeY-50-30, 150, 50, buttonImages, "Valider", RobotoRegularFontF, func() { SaveConfig() })
	}
	// Met à jour l'état du bouton et l'affiche
	accessParticlesButton.Update(screen)
	accessParticlesButton.Draw(screen)

	// // Text input
	// if textInut == nil {
	// 	textInut = newTextInput(10, 300, 100, 50, findImage(ImageList, "text-input-idle.png"), findImage(ImageList, "text-input-hover.png"), findImage(ImageList, "text-input-disabled.png"), true, &config.General.SpawnX, RobotoRegularFontF)
	// }
	// // Update text input
	// textInut.Update(screen)
	// textInut.Draw(screen)

	// // Crée la checkbox si elle n'existe pas
	// if checkbox == nil {
	// 	images := []*ebiten.Image{findImage(ImageList, "checkbox-idle.png"), findImage(ImageList, "checkbox-hover.png"), findImage(ImageList, "checkbox-disabled.png"), findImage(ImageList, "checkbox-checked-idle.png"), findImage(ImageList, "checkbox-unchecked-idle.png"), findImage(ImageList, "checkbox-greyed-disabled.png")}
	// 	checkbox = newCheckbox(10, 70, 50, 50, images, config.General.Debug, true, func() { config.General.Debug = !config.General.Debug })
	// }
	// // Met à jour l'état de la checkbox et l'affiche
	// checkbox.Update(screen)
	// checkbox.Draw(screen)

	// // Create text
	// if texttest == nil {
	// 	texttest = newText(70, 90, "Debug", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	// }
	// // Draw text
	// texttest.Draw(screen)

	// // Create slider
	// if sliderTest == nil {
	// 	sliderTest = newSlider(10, 150, 200, 3, findImage(ImageList, "slider-track-idle.png"), findImage(ImageList, "slider-handle-idle.png"), findImage(ImageList, "slider-handle-hover.png"), findImage(ImageList, "slider-handle-disabled.png"), &config.General.Gravity, 0, 1, true)
	// }
	// // Update slider
	// sliderTest.Update(screen)
	// // Draw slider
	// sliderTest.Draw(screen)

	return nil
}
