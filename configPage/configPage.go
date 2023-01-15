package configPage

import (
	"image/color"
	"log"
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
	// Variable d'initialisation
	itemsCreated = false
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
	spawnRate               *SliderF
	spawnRateText           *Text
	spawnRateValue          *ValueF
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
	rotation                 *SliderF
	rotationValue            *ValueF
	rotationText             *Text
	scaleX                   *SliderF
	scaleXValue              *ValueF
	scaleXText               *Text
	scaleY                   *SliderF
	scaleYValue              *ValueF
	scaleYText               *Text
	opacity                  *SliderF
	opacityValue             *ValueF
	opacityText              *Text
	colorRed                 *SliderF
	colorRedValue            *ValueF
	colorRedText             *Text
	colorGreen               *SliderF
	colorGreenValue          *ValueF
	colorGreenText           *Text
	// Comportement des particules
	behaviorText   *Text
	speedType      *SliderI
	speedTypeValue *ValueI
	speedTypeText  *Text
	gravity        *SliderF
	gravityValue   *ValueF
	gravityText    *Text
	haveLife       *Checkbox
	haveLifeText   *Text
	randomLife     *Checkbox
	randomLifeText *Text
	life           *NumberInput
	lifeText       *Text
	// Gestion de l'écran
	screenManagement         *Text
	marginOutsideScreen      *SliderI
	marginOutsideScreenValue *ValueI
	marginOutsideScreenText  *Text

	accessParticlesButton *Button
)

// Initialisation des variables images
var (
	checkboxImages    = []*ebiten.Image{}
	buttonImages      = []*ebiten.Image{}
	sliderImages      = []*ebiten.Image{}
	numberInputImages = []*ebiten.Image{}
)

// Initialisation des variables de fonts

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

// loadImages charge les images
func loadImages() {
	// Liste des images
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
}

func loadFont() {
	// Liste des polices
	err := loadFontRegular()
	if err != nil {
		log.Fatal(err)
	}
	err = loadFontBold()
	if err != nil {
		log.Fatal(err)
	}
	err = loadFontTitle()
	if err != nil {
		log.Fatal(err)
	}
}

// Creation des éléments de la page de configuration
// La fonction createItems() est appelée une fois au début du programme
// Elle permet de créer tous les éléments de la page de configuration
func createItems() {
	// Titre de configuration
	welcomeTitle = newText(20, 30, "Configuration", PageTitleF, color.RGBA{127, 139, 148, 255})
	// Partie "Configuration de la fenêtre"
	windowConfiguration = newText(20, 60, "Configuration de la fenêtre", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	debugButton = newCheckbox(20, 75, 50, 30, checkboxImages, config.General.Debug, false, func() { config.General.Debug = !config.General.Debug })
	DebugText = newText(100, 95, "Activer le mode debug", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	// Partie "Génération de particules"
	particleGeneration = newText(20, 150, "Génération de particules", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	initNumParticles = newTextInput(20, 165, 100, 30, numberInputImages, false, &config.General.InitNumParticles, RobotoRegularFontF)
	initNumParticlesText = newText(110, 185, "Nombre de particules initiales", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	randomSpawn = newCheckbox(20, 200, 50, 30, checkboxImages, config.General.RandomSpawn, false, func() { config.General.RandomSpawn = !config.General.RandomSpawn })
	randomSpawnText = newText(110, 220, "Activer le spawn aléatoire", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnCenter = newCheckbox(20, 235, 50, 30, checkboxImages, config.General.SpawnCenter, config.General.RandomSpawn, func() { config.General.SpawnCenter = !config.General.SpawnCenter })
	spawnCenterText = newText(110, 255, "Activer le spawn au centre", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnX = newTextInput(20, 270, 100, 30, numberInputImages, config.General.SpawnCenter, &config.General.SpawnX, RobotoRegularFontF)
	spawnXText = newText(110, 290, "Position de spawn X", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnY = newTextInput(20, 305, 100, 30, numberInputImages, config.General.SpawnCenter, &config.General.SpawnY, RobotoRegularFontF)
	spawnYText = newText(110, 325, "Position de spawn Y", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnRate = newSliderFloat(20, 350, 100, 5, sliderImages, &config.General.SpawnRate, 0, 100, false, true)
	spawnRateText = newText(165, 360, "Fréquence de génération", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnRateValue = newValueFloat(125, 360, &config.General.SpawnRate, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnOnAnObject = newCheckbox(20, 375, 50, 30, checkboxImages, config.General.SpawnOnAnObject, false, func() { config.General.SpawnOnAnObject = !config.General.SpawnOnAnObject })
	spawnOnAnObjectText = newText(110, 395, "Activer le spawn sur un objet", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnOnObjectCircle = newCheckbox(40, 410, 50, 30, checkboxImages, config.General.SpawnObject == "circle", !config.General.SpawnOnAnObject, func() { config.General.SpawnObject = "circle" })
	spawnOnObjectCircleText = newText(110, 430, "Cercle", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnOnObjectSquare = newCheckbox(40, 445, 50, 30, checkboxImages, config.General.SpawnObject == "square", !config.General.SpawnOnAnObject, func() { config.General.SpawnObject = "square" })
	spawnOnObjectSquareText = newText(110, 465, "Carré", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	spawnObjectWidth = newTextInput(20, 480, 100, 30, numberInputImages, !config.General.SpawnOnAnObject, &config.General.SpawnObjectWidth, RobotoRegularFontF)
	spawnObjectWidthText = newText(110, 500, "Largeur de l'objet", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	// Partie "Propriétés des particules"
	particulesPropertiesText = newText(20, 550, "Propriétés des particules", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	rotation = newSliderFloat(20, 575, 100, 5, sliderImages, &config.General.Rotation, -3.14159, 3.14159, false, false)
	rotationValue = newValueFloat(125, 585, &config.General.Rotation, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	rotationText = newText(175, 585, "Rotation", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	scaleX = newSliderFloat(20, 600, 100, 5, sliderImages, &config.General.ScaleX, 0, 5, false, true)
	scaleXValue = newValueFloat(125, 610, &config.General.ScaleX, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	scaleXText = newText(175, 610, "ScaleX", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	scaleY = newSliderFloat(20, 625, 100, 5, sliderImages, &config.General.ScaleY, 0, 5, false, true)
	scaleYValue = newValueFloat(125, 635, &config.General.ScaleY, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	scaleYText = newText(175, 635, "ScaleY", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	opacity = newSliderFloat(20, 650, 100, 5, sliderImages, &config.General.Opacity, 0, 1, false, false)
	opacityValue = newValueFloat(125, 660, &config.General.Opacity, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	opacityText = newText(175, 660, "Opacité", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	colorRed = newSliderFloat(20, 675, 100, 5, sliderImages, &config.General.ColorRed, 0, 1, false, false)
	colorRedValue = newValueFloat(125, 685, &config.General.ColorRed, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	colorRedText = newText(175, 685, "Rouge", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	colorGreen = newSliderFloat(20, 700, 100, 5, sliderImages, &config.General.ColorGreen, 0, 1, false, false)
	colorGreenValue = newValueFloat(125, 710, &config.General.ColorGreen, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	colorGreenText = newText(175, 710, "Vert", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	// Partie "Comportement des particules"
	behaviorText = newText(430, 60, "Comportement des particules", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	speedType = newSliderInt(430, 85, 100, 5, sliderImages, &config.General.SpeedType, 0, 3, false)
	speedTypeValue = newValueInt(535, 95, &config.General.SpeedType, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	speedTypeText = newText(585, 95, "Type de vitesse", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	gravity = newSliderFloat(430, 110, 100, 5, sliderImages, &config.General.Gravity, 0, 1, false, false)
	gravityValue = newValueFloat(535, 120, &config.General.Gravity, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	gravityText = newText(585, 120, "Gravité", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	haveLife = newCheckbox(430, 135, 50, 30, checkboxImages, config.General.HaveLife, false, func() { config.General.HaveLife = !config.General.HaveLife })
	haveLifeText = newText(520, 155, "Avec vie", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	randomLife = newCheckbox(430, 170, 50, 30, checkboxImages, config.General.RandomLife, !config.General.HaveLife, func() { config.General.RandomLife = !config.General.RandomLife })
	randomLifeText = newText(520, 190, "Vie aléatoire", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	life = newTextInput(430, 205, 100, 30, numberInputImages, !config.General.HaveLife, &config.General.Life, RobotoRegularFontF)
	lifeText = newText(520, 225, "Durée de vie", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	// Partie "Gestion de l'écran"
	screenManagement = newText(430, 280, "Gestion de l'écran", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	marginOutsideScreen = newSliderInt(430, 305, 100, 5, sliderImages, &config.General.MarginOutsideScreen, -30, 30, false)
	marginOutsideScreenValue = newValueInt(535, 315, &config.General.MarginOutsideScreen, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	marginOutsideScreenText = newText(585, 315, "Marge hors écran", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	accessParticlesButton = newButton(config.General.WindowSizeX-150-30, config.General.WindowSizeY-50-30, 150, 50, buttonImages, "Valider", RobotoRegularFontF, func() { SaveConfig() })
}

// UpdateConfigPage est la fonction qui est appelée pour mettre à jour la page de configuration
func UpdateConfigPage(screen *ebiten.Image) error {
	if !itemsCreated {
		loadImages()
		loadFont()
		createItems()
		itemsCreated = true
	}
	// EbitenImages

	// Titre de configuration
	welcomeTitle.Draw(screen)

	// Titre de configuration de la fenêtre
	windowConfiguration.Draw(screen)

	// Checkbox debug
	debugButton.Update(screen)
	debugButton.Draw(screen)
	DebugText.Draw(screen)

	// Titre de génération de particules
	particleGeneration.Draw(screen)

	// Nombre de particules initiales
	initNumParticles.Update(screen)
	initNumParticles.Draw(screen)
	initNumParticlesText.Draw(screen)

	// Checkbox random spawn
	randomSpawn.Update(screen)
	randomSpawn.Draw(screen)
	randomSpawnText.Draw(screen)

	// Checkbox spawn center
	spawnCenter.disabled = config.General.RandomSpawn
	spawnCenter.Update(screen)
	spawnCenter.Draw(screen)
	spawnCenterText.Draw(screen)

	// Input de la position de spawnX
	spawnX.disabled = config.General.SpawnCenter
	spawnX.Update(screen)
	spawnX.Draw(screen)
	spawnXText.Draw(screen)

	// Input de la position de spawnY
	spawnY.disabled = config.General.SpawnCenter
	spawnY.Update(screen)
	spawnY.Draw(screen)
	spawnYText.Draw(screen)

	// Input SpawnRate Slider
	spawnRate.Update(screen)
	spawnRate.Draw(screen)
	spawnRateText.Draw(screen)
	spawnRateValue.Update(screen)
	spawnRateValue.Draw(screen)

	// Checkbox SpawnOnAnObject
	spawnOnAnObject.Update(screen)
	spawnOnAnObject.Draw(screen)
	spawnOnAnObjectText.Draw(screen)

	// Checkbox SpawnObject = "circle"
	spawnOnObjectCircle.checked = config.General.SpawnObject == "circle"
	spawnOnObjectCircle.Update(screen)
	spawnOnObjectCircle.Draw(screen)
	spawnOnObjectCircleText.Draw(screen)

	// Checkbox SpawnObject = "square"
	spawnOnObjectSquare.checked = config.General.SpawnObject == "square"
	spawnOnObjectSquare.Update(screen)
	spawnOnObjectSquare.Draw(screen)
	spawnOnObjectSquareText.Draw(screen)
	// (Vérifie que SpawnOnAnObject est activé)
	spawnOnObjectCircle.disabled, spawnOnObjectSquare.disabled = !config.General.SpawnOnAnObject, !config.General.SpawnOnAnObject

	// Input spawnObjectWidth number
	spawnObjectWidth.disabled = !config.General.SpawnOnAnObject
	spawnObjectWidth.Update(screen)
	spawnObjectWidth.Draw(screen)
	spawnObjectWidthText.Draw(screen)

	// Titre particulesPropertiesText
	particulesPropertiesText.Draw(screen)

	// Slider rotation
	rotation.Update(screen)
	rotation.Draw(screen)

	rotationValue.Update(screen)
	rotationValue.Draw(screen)

	rotationText.Draw(screen)

	// Slider ScaleX
	scaleX.Update(screen)
	scaleX.Draw(screen)

	scaleXValue.Update(screen)
	scaleXValue.Draw(screen)

	scaleXText.Draw(screen)
	// Slider ScaleY
	scaleY.Update(screen)
	scaleY.Draw(screen)

	scaleYValue.Update(screen)
	scaleYValue.Draw(screen)

	scaleYText.Draw(screen)

	// Opacity
	opacity.Update(screen)
	opacity.Draw(screen)

	opacityValue.Update(screen)
	opacityValue.Draw(screen)

	opacityText.Draw(screen)

	// ColorRed
	colorRed.Update(screen)
	colorRed.Draw(screen)

	colorRedValue.Update(screen)
	colorRedValue.Draw(screen)

	colorRedText.Draw(screen)

	// ColorGreen
	colorGreen.Update(screen)
	colorGreen.Draw(screen)

	colorGreenValue.Update(screen)
	colorGreenValue.Draw(screen)

	colorGreenText.Draw(screen)

	// Comportement des particules
	// Titre
	behaviorText.Draw(screen)

	// SpeedType (sliderI)
	speedType.Update(screen)
	speedType.Draw(screen)

	speedTypeValue.Update(screen)
	speedTypeValue.Draw(screen)

	speedTypeText.Draw(screen)

	// Gravity (sliderF)
	gravity.Update(screen)
	gravity.Draw(screen)

	gravityValue.Update(screen)
	gravityValue.Draw(screen)

	gravityText.Draw(screen)

	// haveLife Checkbox
	haveLife.Update(screen)
	haveLife.Draw(screen)

	haveLifeText.Draw(screen)

	// randomLife Checkbox
	randomLife.disabled = !config.General.HaveLife
	randomLife.Update(screen)
	randomLife.Draw(screen)

	randomLifeText.Draw(screen)

	// Life numberInput
	life.disabled = (config.General.HaveLife && config.General.RandomLife) || !config.General.HaveLife
	life.Update(screen)
	life.Draw(screen)

	lifeText.Draw(screen)

	// screenManagement title
	screenManagement.Draw(screen)

	// marginOutsideScreen slider
	marginOutsideScreen.Update(screen)
	marginOutsideScreen.Draw(screen)

	marginOutsideScreenValue.Update(screen)
	marginOutsideScreenValue.Draw(screen)

	marginOutsideScreenText.Draw(screen)

	// Met à jour l'état du bouton et l'affiche
	accessParticlesButton.Update(screen)
	accessParticlesButton.Draw(screen)

	return nil
}
