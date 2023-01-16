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
	debugText               *Text
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
	colorBlue                *SliderF
	colorBlueValue           *ValueF
	colorBlueText            *Text
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
	// Changements dynamiques
	dynamicChangesText                   *Text
	changeColorAccordingToTitle          *Text
	changeColorAccordingToNothing        *Checkbox
	changeColorAccordingToNothingText    *Text
	changeColorAccordingToPosition       *Checkbox
	changeColorAccordingToPositionText   *Text
	changeColorAccordingToLife           *Checkbox
	changeColorAccordingToLifeText       *Text
	minColorRed                          *SliderF
	minColorRedValue                     *ValueF
	minColorRedText                      *Text
	minColorGreen                        *SliderF
	minColorGreenValue                   *ValueF
	minColorGreenText                    *Text
	minColorBlue                         *SliderF
	minColorBlueValue                    *ValueF
	minColorBlueText                     *Text
	maxColorRed                          *SliderF
	maxColorRedValue                     *ValueF
	maxColorRedText                      *Text
	maxColorGreen                        *SliderF
	maxColorGreenValue                   *ValueF
	maxColorGreenText                    *Text
	maxColorBlue                         *SliderF
	maxColorBlueValue                    *ValueF
	maxColorBlueText                     *Text
	changeScaleAccordingToTitle          *Text
	changeScaleAccordingToNothing        *Checkbox
	changeScaleAccordingToNothingText    *Text
	changeScaleAccordingToPosition       *Checkbox
	changeScaleAccordingToPositionText   *Text
	changeScaleAccordingToLife           *Checkbox
	changeScaleAccordingToLifeText       *Text
	changeOpacityAccordingToTitle        *Text
	changeOpacityAccordingToNothing      *Checkbox
	changeOpacityAccordingToNothingText  *Text
	changeOpacityAccordingToPosition     *Checkbox
	changeOpacityAccordingToPositionText *Text
	changeOpacityAccordingToLife         *Checkbox
	changeOpacityAccordingToLifeText     *Text

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
	debugText = newText(100, 95, "Activer le mode debug", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
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
	colorBlue = newSliderFloat(20, 725, 100, 5, sliderImages, &config.General.ColorBlue, 0, 1, false, false)
	colorBlueValue = newValueFloat(125, 735, &config.General.ColorBlue, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	colorBlueText = newText(175, 735, "Bleu", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
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
	dynamicChangesText = newText(430, 360, "Changements dynamiques", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	changeColorAccordingToTitle = newText(430, 385, "Changer la couleur en fonction de :", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeColorAccordingToNothing = newCheckbox(430, 395, 50, 30, checkboxImages, config.General.ChangeColorAccordingTo == 0, false, func() {
		if changeColorAccordingToNothing.checked {
			config.General.ChangeColorAccordingTo = 0
		}
	})
	changeColorAccordingToNothingText = newText(520, 415, "Aucun", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeColorAccordingToPosition = newCheckbox(430, 430, 50, 30, checkboxImages, config.General.ChangeColorAccordingTo == 1, false, func() {
		if changeColorAccordingToPosition.checked {
			config.General.ChangeColorAccordingTo = 1
		}
	})
	changeColorAccordingToPositionText = newText(520, 450, "En fonction de la position", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeColorAccordingToLife = newCheckbox(430, 465, 50, 30, checkboxImages, config.General.ChangeColorAccordingTo == 2, false, func() {
		if changeColorAccordingToLife.checked {
			config.General.ChangeColorAccordingTo = 2
		}
	})
	changeColorAccordingToLifeText = newText(520, 485, "En fonction de la vie", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	minColorRed = newSliderFloat(430, 510, 100, 5, sliderImages, &config.General.MinColorRed, 0, 1, false, false)
	minColorRedValue = newValueFloat(535, 520, &config.General.MinColorRed, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	minColorRedText = newText(585, 520, "Couleur minimale (Rouge)", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	minColorGreen = newSliderFloat(430, 535, 100, 5, sliderImages, &config.General.MinColorGreen, 0, 1, false, false)
	minColorGreenValue = newValueFloat(535, 545, &config.General.MinColorGreen, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	minColorGreenText = newText(585, 545, "Couleur minimale (Vert)", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	minColorBlue = newSliderFloat(430, 560, 100, 5, sliderImages, &config.General.MinColorBlue, 0, 1, false, false)
	minColorBlueValue = newValueFloat(535, 570, &config.General.MinColorBlue, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	minColorBlueText = newText(585, 570, "Couleur minimale (Bleu)", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	maxColorRed = newSliderFloat(430, 595, 100, 5, sliderImages, &config.General.MaxColorRed, 0, 1, false, false)
	maxColorRedValue = newValueFloat(535, 605, &config.General.MaxColorRed, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	maxColorRedText = newText(585, 605, "Couleur maximale (Rouge)", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	maxColorGreen = newSliderFloat(430, 620, 100, 5, sliderImages, &config.General.MaxColorGreen, 0, 1, false, false)
	maxColorGreenValue = newValueFloat(535, 630, &config.General.MaxColorGreen, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	maxColorGreenText = newText(585, 630, "Couleur maximale (Vert)", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	maxColorBlue = newSliderFloat(430, 645, 100, 5, sliderImages, &config.General.MaxColorBlue, 0, 1, false, false)
	maxColorBlueValue = newValueFloat(535, 655, &config.General.MaxColorBlue, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	maxColorBlueText = newText(585, 655, "Couleur maximale (Bleu)", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeScaleAccordingToTitle = newText(430, 690, "Changer l'échelle en fonction de :", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeScaleAccordingToNothing = newCheckbox(430, 700, 50, 30, checkboxImages, config.General.ChangeScaleAccordingTo == 0, false, func() {
		if changeScaleAccordingToNothing.checked {
			config.General.ChangeScaleAccordingTo = 0
		}
	})
	changeScaleAccordingToNothingText = newText(520, 720, "Aucun", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeScaleAccordingToPosition = newCheckbox(430, 735, 50, 30, checkboxImages, config.General.ChangeScaleAccordingTo == 1, false, func() {
		if changeScaleAccordingToPosition.checked {
			config.General.ChangeScaleAccordingTo = 1
		}
	})
	changeScaleAccordingToPositionText = newText(520, 755, "En fonction de la position", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeScaleAccordingToLife = newCheckbox(430, 770, 50, 30, checkboxImages, config.General.ChangeScaleAccordingTo == 2, false, func() {
		if changeScaleAccordingToLife.checked {
			config.General.ChangeScaleAccordingTo = 2
		}
	})
	changeScaleAccordingToLifeText = newText(520, 790, "En fonction de la vie", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeOpacityAccordingToTitle = newText(430, 835, "Changer l'opacité en fonction de :", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeOpacityAccordingToNothing = newCheckbox(430, 845, 50, 30, checkboxImages, config.General.ChangeOpacityAccordingTo == 0, false, func() {
		if changeOpacityAccordingToNothing.checked {
			config.General.ChangeOpacityAccordingTo = 0
		}
	})
	changeOpacityAccordingToNothingText = newText(520, 865, "Aucun", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeOpacityAccordingToPosition = newCheckbox(430, 880, 50, 30, checkboxImages, config.General.ChangeOpacityAccordingTo == 1, false, func() {
		if changeOpacityAccordingToPosition.checked {
			config.General.ChangeOpacityAccordingTo = 1
		}
	})
	changeOpacityAccordingToPositionText = newText(520, 900, "En fonction de la position", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	changeOpacityAccordingToLife = newCheckbox(430, 915, 50, 30, checkboxImages, config.General.ChangeOpacityAccordingTo == 2, false, func() {
		if changeOpacityAccordingToLife.checked {
			config.General.ChangeOpacityAccordingTo = 2
		}
	})
	changeOpacityAccordingToLifeText = newText(520, 935, "En fonction de la vie", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})

}

var (
	ScrollY int
)

// UpdateConfigPage est la fonction qui est appelée pour mettre à jour la page de configuration
func UpdateConfigPage(screen *ebiten.Image) error {
	if !itemsCreated {
		loadImages()
		loadFont()
		createItems()
		itemsCreated = true
	}

	// Ajouter ScrollY à la page de configuration

	// Titre de configuration
	welcomeTitle.y += int(ScrollY)
	welcomeTitle.Draw(screen)

	// Titre de configuration de la fenêtre
	windowConfiguration.y += int(ScrollY)
	windowConfiguration.Draw(screen)

	// Checkbox debug
	debugButton.Update(screen)
	debugButton.y += int(ScrollY)
	debugButton.Draw(screen)
	debugText.y += int(ScrollY)
	debugText.Draw(screen)

	// Titre de génération de particules
	particleGeneration.y += int(ScrollY)
	particleGeneration.Draw(screen)

	// Nombre de particules initiales
	initNumParticles.Update(screen)
	initNumParticles.y += int(ScrollY)
	initNumParticles.Draw(screen)
	initNumParticlesText.y += int(ScrollY)
	initNumParticlesText.Draw(screen)

	// Checkbox random spawn
	randomSpawn.Update(screen)
	randomSpawn.y += int(ScrollY)
	randomSpawn.Draw(screen)
	randomSpawnText.y += int(ScrollY)
	randomSpawnText.Draw(screen)

	// Checkbox spawn center
	spawnCenter.disabled = config.General.RandomSpawn
	spawnCenter.Update(screen)
	spawnCenter.y += int(ScrollY)
	spawnCenter.Draw(screen)
	spawnCenterText.y += int(ScrollY)
	spawnCenterText.Draw(screen)

	// Input de la position de spawnX
	spawnX.disabled = config.General.SpawnCenter
	spawnX.Update(screen)
	spawnX.y += int(ScrollY)
	spawnX.Draw(screen)
	spawnXText.y += int(ScrollY)
	spawnXText.Draw(screen)

	// Input de la position de spawnY
	spawnY.disabled = config.General.SpawnCenter
	spawnY.Update(screen)
	spawnY.y += int(ScrollY)
	spawnY.Draw(screen)
	spawnYText.y += int(ScrollY)
	spawnYText.Draw(screen)

	// Input SpawnRate Slider
	spawnRate.Update(screen)
	spawnRate.y += int(ScrollY)
	spawnRate.Draw(screen)
	spawnRateText.y += int(ScrollY)
	spawnRateText.Draw(screen)
	spawnRateValue.Update(screen)
	spawnRateValue.y += int(ScrollY)
	spawnRateValue.Draw(screen)

	// Checkbox SpawnOnAnObject
	spawnOnAnObject.Update(screen)
	spawnOnAnObject.y += int(ScrollY)
	spawnOnAnObject.Draw(screen)
	spawnOnAnObjectText.y += int(ScrollY)
	spawnOnAnObjectText.Draw(screen)

	// Checkbox SpawnObject = "circle"
	spawnOnObjectCircle.checked = config.General.SpawnObject == "circle"
	spawnOnObjectCircle.Update(screen)
	spawnOnObjectCircle.y += int(ScrollY)
	spawnOnObjectCircle.Draw(screen)
	spawnOnObjectCircleText.y += int(ScrollY)
	spawnOnObjectCircleText.Draw(screen)

	// Checkbox SpawnObject = "square"
	spawnOnObjectSquare.checked = config.General.SpawnObject == "square"
	spawnOnObjectSquare.Update(screen)
	spawnOnObjectSquare.y += int(ScrollY)
	spawnOnObjectSquare.Draw(screen)
	spawnOnObjectSquareText.y += int(ScrollY)
	spawnOnObjectSquareText.Draw(screen)
	// (Vérifie que SpawnOnAnObject est activé)
	spawnOnObjectCircle.disabled, spawnOnObjectSquare.disabled = !config.General.SpawnOnAnObject, !config.General.SpawnOnAnObject

	// Input spawnObjectWidth number
	spawnObjectWidth.disabled = !config.General.SpawnOnAnObject
	spawnObjectWidth.Update(screen)
	spawnObjectWidth.y += int(ScrollY)
	spawnObjectWidth.Draw(screen)
	spawnObjectWidthText.y += int(ScrollY)
	spawnObjectWidthText.Draw(screen)

	// Titre particulesPropertiesText
	particulesPropertiesText.y += int(ScrollY)
	particulesPropertiesText.Draw(screen)

	// Slider rotation
	rotation.Update(screen)
	rotation.y += int(ScrollY)
	rotation.Draw(screen)

	rotationValue.Update(screen)
	rotationValue.y += int(ScrollY)
	rotationValue.Draw(screen)

	rotationText.y += int(ScrollY)
	rotationText.Draw(screen)

	// Slider ScaleX
	scaleX.Update(screen)
	scaleX.y += int(ScrollY)
	scaleX.Draw(screen)

	scaleXValue.Update(screen)
	scaleXValue.y += int(ScrollY)
	scaleXValue.Draw(screen)

	scaleXText.y += int(ScrollY)
	scaleXText.Draw(screen)
	// Slider ScaleY
	scaleY.Update(screen)
	scaleY.y += int(ScrollY)
	scaleY.Draw(screen)

	scaleYValue.Update(screen)
	scaleYValue.y += int(ScrollY)
	scaleYValue.Draw(screen)

	scaleYText.y += int(ScrollY)
	scaleYText.Draw(screen)

	// Opacity
	opacity.Update(screen)
	opacity.y += int(ScrollY)
	opacity.Draw(screen)

	opacityValue.Update(screen)
	opacityValue.y += int(ScrollY)
	opacityValue.Draw(screen)

	opacityText.y += int(ScrollY)
	opacityText.Draw(screen)

	// ColorRed
	colorRed.Update(screen)
	colorRed.y += int(ScrollY)
	colorRed.Draw(screen)

	colorRedValue.Update(screen)
	colorRedValue.y += int(ScrollY)
	colorRedValue.Draw(screen)

	colorRedText.y += int(ScrollY)
	colorRedText.Draw(screen)

	// ColorGreen
	colorGreen.Update(screen)
	colorGreen.y += int(ScrollY)
	colorGreen.Draw(screen)

	colorGreenValue.Update(screen)
	colorGreenValue.y += int(ScrollY)
	colorGreenValue.Draw(screen)

	colorGreenText.y += int(ScrollY)
	colorGreenText.Draw(screen)

	// colorBlue
	colorBlue.Update(screen)
	colorBlue.y += int(ScrollY)
	colorBlue.Draw(screen)

	colorBlueValue.Update(screen)
	colorBlueValue.y += int(ScrollY)
	colorBlueValue.Draw(screen)

	colorBlueText.y += int(ScrollY)
	colorBlueText.Draw(screen)

	// Comportement des particules
	// Titre
	behaviorText.y += int(ScrollY)
	behaviorText.Draw(screen)

	// SpeedType (sliderI)
	speedType.Update(screen)
	speedType.y += int(ScrollY)
	speedType.Draw(screen)

	speedTypeValue.Update(screen)
	speedTypeValue.y += int(ScrollY)
	speedTypeValue.Draw(screen)

	speedTypeText.y += int(ScrollY)
	speedTypeText.Draw(screen)

	// Gravity (sliderF)
	gravity.Update(screen)
	gravity.y += int(ScrollY)
	gravity.Draw(screen)

	gravityValue.Update(screen)
	gravityValue.y += int(ScrollY)
	gravityValue.Draw(screen)

	gravityText.y += int(ScrollY)
	gravityText.Draw(screen)

	// haveLife Checkbox
	haveLife.Update(screen)
	haveLife.y += int(ScrollY)
	haveLife.Draw(screen)

	haveLifeText.y += int(ScrollY)
	haveLifeText.Draw(screen)

	// randomLife Checkbox
	randomLife.disabled = !config.General.HaveLife
	randomLife.Update(screen)
	randomLife.y += int(ScrollY)
	randomLife.Draw(screen)

	randomLifeText.y += int(ScrollY)
	randomLifeText.Draw(screen)

	// Life numberInput
	life.disabled = (config.General.HaveLife && config.General.RandomLife) || !config.General.HaveLife
	life.Update(screen)
	life.y += int(ScrollY)
	life.Draw(screen)

	lifeText.y += int(ScrollY)
	lifeText.Draw(screen)

	// screenManagement title
	screenManagement.y += int(ScrollY)
	screenManagement.Draw(screen)

	// marginOutsideScreen slider
	marginOutsideScreen.Update(screen)
	marginOutsideScreen.y += int(ScrollY)
	marginOutsideScreen.Draw(screen)

	marginOutsideScreenValue.Update(screen)
	marginOutsideScreenValue.y += int(ScrollY)
	marginOutsideScreenValue.Draw(screen)

	marginOutsideScreenText.y += int(ScrollY)
	marginOutsideScreenText.Draw(screen)

	// dynamicChanges title
	dynamicChangesText.y += int(ScrollY)
	dynamicChangesText.Draw(screen)

	// changeColorAccordingToTitle
	changeColorAccordingToTitle.y += int(ScrollY)
	changeColorAccordingToTitle.Draw(screen)

	// changeColorAccordingToNothing
	changeColorAccordingToNothing.checked = config.General.ChangeColorAccordingTo == 0
	changeColorAccordingToNothing.Update(screen)
	changeColorAccordingToNothing.y += int(ScrollY)
	changeColorAccordingToNothing.Draw(screen)

	changeColorAccordingToNothingText.y += int(ScrollY)
	changeColorAccordingToNothingText.Draw(screen)

	// changeColorAccordingToPosition
	changeColorAccordingToPosition.checked = config.General.ChangeColorAccordingTo == 1
	changeColorAccordingToPosition.Update(screen)
	changeColorAccordingToPosition.y += int(ScrollY)
	changeColorAccordingToPosition.Draw(screen)

	changeColorAccordingToPositionText.y += int(ScrollY)
	changeColorAccordingToPositionText.Draw(screen)

	// changeColorAccordingToLife
	changeColorAccordingToLife.checked = config.General.ChangeColorAccordingTo == 2
	changeColorAccordingToLife.disabled = !config.General.HaveLife
	changeColorAccordingToLife.Update(screen)
	changeColorAccordingToLife.y += int(ScrollY)
	changeColorAccordingToLife.Draw(screen)
	if config.General.ChangeColorAccordingTo == 2 && !config.General.HaveLife {
		config.General.ChangeColorAccordingTo = 0
	}

	changeColorAccordingToLifeText.y += int(ScrollY)
	changeColorAccordingToLifeText.Draw(screen)

	// minColorRed
	minColorRed.Update(screen)
	minColorRed.y += int(ScrollY)
	minColorRed.Draw(screen)
	minColorRedValue.Update(screen)
	minColorRedValue.y += int(ScrollY)
	minColorRedValue.Draw(screen)
	minColorRedText.y += int(ScrollY)
	minColorRedText.Draw(screen)

	// minColorGreen
	minColorGreen.Update(screen)
	minColorGreen.y += int(ScrollY)
	minColorGreen.Draw(screen)
	minColorGreenValue.Update(screen)
	minColorGreenValue.y += int(ScrollY)
	minColorGreenValue.Draw(screen)
	minColorGreenText.y += int(ScrollY)
	minColorGreenText.Draw(screen)

	// minColorBlue
	minColorBlue.Update(screen)
	minColorBlue.y += int(ScrollY)
	minColorBlue.Draw(screen)
	minColorBlueValue.Update(screen)
	minColorBlueValue.y += int(ScrollY)
	minColorBlueValue.Draw(screen)
	minColorBlueText.y += int(ScrollY)
	minColorBlueText.Draw(screen)

	// maxColorRed
	maxColorRed.Update(screen)
	maxColorRed.y += int(ScrollY)
	maxColorRed.Draw(screen)
	maxColorRedValue.Update(screen)
	maxColorRedValue.y += int(ScrollY)
	maxColorRedValue.Draw(screen)
	maxColorRedText.y += int(ScrollY)
	maxColorRedText.Draw(screen)

	// maxColorGreen
	maxColorGreen.Update(screen)
	maxColorGreen.y += int(ScrollY)
	maxColorGreen.Draw(screen)
	maxColorGreenValue.Update(screen)
	maxColorGreenValue.y += int(ScrollY)
	maxColorGreenValue.Draw(screen)
	maxColorGreenText.y += int(ScrollY)
	maxColorGreenText.Draw(screen)

	// maxColorBlue
	maxColorBlue.Update(screen)
	maxColorBlue.y += int(ScrollY)
	maxColorBlue.Draw(screen)
	maxColorBlueValue.Update(screen)
	maxColorBlueValue.y += int(ScrollY)
	maxColorBlueValue.Draw(screen)
	maxColorBlueText.y += int(ScrollY)
	maxColorBlueText.Draw(screen)

	// changeScaleAccordingToTitle
	changeScaleAccordingToTitle.y += int(ScrollY)
	changeScaleAccordingToTitle.Draw(screen)

	// changeScaleAccordingToNothing
	changeScaleAccordingToNothing.checked = config.General.ChangeScaleAccordingTo == 0
	changeScaleAccordingToNothing.Update(screen)
	changeScaleAccordingToNothing.y += int(ScrollY)
	changeScaleAccordingToNothing.Draw(screen)

	changeScaleAccordingToNothingText.y += int(ScrollY)
	changeScaleAccordingToNothingText.Draw(screen)

	// changeScaleAccordingToPosition
	changeScaleAccordingToPosition.checked = config.General.ChangeScaleAccordingTo == 1
	changeScaleAccordingToPosition.Update(screen)
	changeScaleAccordingToPosition.y += int(ScrollY)
	changeScaleAccordingToPosition.Draw(screen)

	changeScaleAccordingToPositionText.y += int(ScrollY)
	changeScaleAccordingToPositionText.Draw(screen)

	// changeScaleAccordingToLife
	changeScaleAccordingToLife.checked = config.General.ChangeScaleAccordingTo == 2
	changeScaleAccordingToLife.disabled = !config.General.HaveLife
	changeScaleAccordingToLife.Update(screen)
	changeScaleAccordingToLife.y += int(ScrollY)
	changeScaleAccordingToLife.Draw(screen)

	changeScaleAccordingToLifeText.y += int(ScrollY)
	changeScaleAccordingToLifeText.Draw(screen)

	// changeOpacityAccordingToTitle
	changeOpacityAccordingToTitle.y += int(ScrollY)
	changeOpacityAccordingToTitle.Draw(screen)

	// changeOpacityAccordingToNothing
	changeOpacityAccordingToNothing.checked = config.General.ChangeOpacityAccordingTo == 0
	changeOpacityAccordingToNothing.Update(screen)
	changeOpacityAccordingToNothing.y += int(ScrollY)
	changeOpacityAccordingToNothing.Draw(screen)

	changeOpacityAccordingToNothingText.y += int(ScrollY)
	changeOpacityAccordingToNothingText.Draw(screen)

	// changeOpacityAccordingToPosition
	changeOpacityAccordingToPosition.checked = config.General.ChangeOpacityAccordingTo == 1
	changeOpacityAccordingToPosition.Update(screen)
	changeOpacityAccordingToPosition.y += int(ScrollY)
	changeOpacityAccordingToPosition.Draw(screen)

	changeOpacityAccordingToPositionText.y += int(ScrollY)
	changeOpacityAccordingToPositionText.Draw(screen)

	// changeOpacityAccordingToLife
	changeOpacityAccordingToLife.checked = config.General.ChangeOpacityAccordingTo == 2
	changeOpacityAccordingToLife.disabled = !config.General.HaveLife
	changeOpacityAccordingToLife.Update(screen)
	changeOpacityAccordingToLife.y += int(ScrollY)
	changeOpacityAccordingToLife.Draw(screen)

	changeOpacityAccordingToLifeText.y += int(ScrollY)
	changeOpacityAccordingToLifeText.Draw(screen)

	// Met à jour l'état du bouton et l'affiche
	accessParticlesButton.Update(screen)
	accessParticlesButton.Draw(screen)

	return nil
}
