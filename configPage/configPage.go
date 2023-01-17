package configPage

import (
	"image/color"
	"log"
	"os"
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
	// Partie collisions
	collisionsText             *Text
	collision                  *Checkbox
	collisionText              *Text
	collisionAmongParticle     *Checkbox
	collisionAmongParticleText *Text
	collisionWithWall          *Checkbox
	collisionWithWallText      *Text
	// Interactions
	interactionsText *Text
	interaction      *Checkbox
	interactionText  *Text
	followMouse      *Checkbox
	followMouseText  *Text
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

	// Boutons
	PlayButton            *Button
	accessParticlesButton *Button
	leaveGamebutton       *Button
)

// Initialisation des variables images
var (
	checkboxImages    = []*ebiten.Image{}
	buttonImages      = []*ebiten.Image{}
	sliderImages      = []*ebiten.Image{}
	numberInputImages = []*ebiten.Image{}
)

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
	scaleX = newSliderFloat(20, 600, 100, 5, sliderImages, &config.General.ScaleX, 0, 5, false, false)
	scaleXValue = newValueFloat(125, 610, &config.General.ScaleX, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	scaleXText = newText(175, 610, "ScaleX", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	scaleY = newSliderFloat(20, 625, 100, 5, sliderImages, &config.General.ScaleY, 0, 5, false, false)
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
	// Partie Collisions
	collisionsText = newText(20, 785, "Collisions", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	collision = newCheckbox(20, 810, 50, 30, checkboxImages, config.General.Collision, false, func() { config.General.Collision = !config.General.Collision })
	collisionText = newText(110, 830, "Activer les collisions", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	collisionAmongParticle = newCheckbox(20, 845, 50, 30, checkboxImages, config.General.CollisionAmongParticle, !config.General.Collision, func() { config.General.CollisionAmongParticle = !config.General.CollisionAmongParticle })
	collisionAmongParticleText = newText(110, 865, "Entre les particules", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	collisionWithWall = newCheckbox(20, 880, 50, 30, checkboxImages, config.General.CollisionWithWall, !config.General.Collision, func() { config.General.CollisionWithWall = !config.General.CollisionWithWall })
	collisionWithWallText = newText(110, 900, "Avec les murs", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	// Interactions
	interactionsText = newText(20, 950, "Interactions", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	interaction = newCheckbox(20, 975, 50, 30, checkboxImages, config.General.Interaction, false, func() { config.General.Interaction = !config.General.Interaction })
	interactionText = newText(110, 995, "Activer les interactions", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	followMouse = newCheckbox(20, 1010, 50, 30, checkboxImages, config.General.FollowMouse, !config.General.Interaction, func() { config.General.FollowMouse = !config.General.FollowMouse })
	followMouseText = newText(110, 1030, "Suivre la souris", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	// Partie "Comportement des particules"
	behaviorText = newText(430, 60, "Comportement des particules", RobotoBoldFontF, color.RGBA{127, 139, 148, 255})
	speedType = newSliderInt(430, 85, 100, 5, sliderImages, &config.General.SpeedType, 0, 3, false)
	speedTypeValue = newValueInt(535, 95, &config.General.SpeedType, RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	speedTypeText = newText(585, 95, "Type de vitesse", RobotoRegularFontF, color.RGBA{127, 139, 148, 255})
	gravity = newSliderFloat(430, 110, 100, 5, sliderImages, &config.General.Gravity, -1, 1, false, false)
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
	PlayButton = newButton(config.General.WindowSizeX-150-50, config.General.WindowSizeY-50-130, 170, 50, buttonImages, "Jouer", RobotoRegularFontF, func() {})
	accessParticlesButton = newButton(config.General.WindowSizeX-150-50, config.General.WindowSizeY-50-80, 170, 50, buttonImages, "Sauvegarder", RobotoRegularFontF, func() { SaveConfig() })
	leaveGamebutton = newButton(config.General.WindowSizeX-150-50, config.General.WindowSizeY-50-30, 170, 50, buttonImages, "Quitter", RobotoRegularFontF, func() { os.Exit(0) })
}

// Cette variable est utilisée pour le Scroll vertical
// Elle est ajoutée aux composants dans leurs fonctions Draw
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

	// Titre de configuration
	welcomeTitle.update(screen)

	// Titre de configuration de la fenêtre
	windowConfiguration.update(screen)

	// Checkbox debug
	debugButton.update(screen)
	debugText.update(screen)

	// Paartie génération de particules
	// Titre de génération de particules
	particleGeneration.update(screen)

	// Nombre de particules initiales
	initNumParticles.update(screen)
	initNumParticlesText.update(screen)

	// Checkbox random spawn
	randomSpawn.update(screen)
	randomSpawnText.update(screen)

	// Checkbox spawn center
	spawnCenter.disabled = config.General.RandomSpawn
	spawnCenter.update(screen)
	spawnCenterText.update(screen)

	// Input de la position de spawnX
	spawnX.disabled = config.General.SpawnCenter
	spawnX.update(screen)
	spawnXText.update(screen)

	// Input de la position de spawnY
	spawnY.disabled = config.General.SpawnCenter
	spawnY.update(screen)
	spawnYText.update(screen)

	// Input SpawnRate Slider
	spawnRate.update(screen)
	spawnRateText.update(screen)
	spawnRateValue.update(screen)

	// Checkbox SpawnOnAnObject
	spawnOnAnObject.update(screen)
	spawnOnAnObjectText.update(screen)

	// Checkbox SpawnObject = "circle"
	spawnOnObjectCircle.checked = config.General.SpawnObject == "circle"
	spawnOnObjectCircle.update(screen)
	spawnOnObjectCircleText.update(screen)

	// Checkbox SpawnObject = "square"
	spawnOnObjectSquare.checked = config.General.SpawnObject == "square"
	spawnOnObjectSquare.update(screen)
	spawnOnObjectSquareText.update(screen)
	// (Vérifie que SpawnOnAnObject est activé)
	spawnOnObjectCircle.disabled, spawnOnObjectSquare.disabled = !config.General.SpawnOnAnObject, !config.General.SpawnOnAnObject

	// Input spawnObjectWidth number
	spawnObjectWidth.disabled = !config.General.SpawnOnAnObject
	spawnObjectWidth.update(screen)
	spawnObjectWidthText.update(screen)

	// Titre particulesPropertiesText
	particulesPropertiesText.update(screen)

	// Slider rotation
	rotation.update(screen)
	rotationValue.update(screen)
	rotationText.update(screen)

	// Slider ScaleX
	scaleX.update(screen)
	scaleXValue.update(screen)
	scaleXText.update(screen)

	// Slider ScaleY
	scaleY.update(screen)
	scaleYValue.update(screen)
	scaleYText.update(screen)

	// Opacity
	opacity.update(screen)
	opacityValue.update(screen)
	opacityText.update(screen)

	// ColorRed
	colorRed.update(screen)
	colorRedValue.update(screen)
	colorRedText.update(screen)

	// ColorGreen
	colorGreen.update(screen)
	colorGreenValue.update(screen)
	colorGreenText.update(screen)

	// colorBlue
	colorBlue.update(screen)
	colorBlueValue.update(screen)
	colorBlueText.update(screen)

	// Partie collisions
	// collisionsText
	collisionsText.update(screen)

	// collision
	collision.update(screen)
	collisionText.update(screen)

	if !collision.checked {
		collisionAmongParticle.disabled = true
		collisionWithWall.disabled = true
	} else {
		collisionAmongParticle.disabled = false
		collisionWithWall.disabled = false
	}

	// CollisionAmongParticle
	collisionAmongParticle.update(screen)
	collisionAmongParticleText.update(screen)

	// CollisionWithWall
	collisionWithWall.update(screen)
	collisionWithWallText.update(screen)

	// Partie interaction
	// interactionText
	interactionsText.update(screen)

	// interaction
	interaction.update(screen)
	interactionText.update(screen)

	if !interaction.checked || randomSpawn.checked {
		followMouse.disabled = true
	} else {
		followMouse.disabled = false
	}

	// FollowMouse
	followMouse.update(screen)
	followMouseText.update(screen)

	// Comportement des particules
	// Titre
	behaviorText.update(screen)

	// SpeedType (sliderI)
	speedType.update(screen)
	speedTypeValue.update(screen)
	speedTypeText.update(screen)

	// Gravity (sliderF)
	gravity.update(screen)
	gravityValue.update(screen)
	gravityText.update(screen)

	// haveLife Checkbox
	haveLife.update(screen)
	haveLifeText.update(screen)

	// randomLife Checkbox
	randomLife.disabled = !config.General.HaveLife
	randomLife.update(screen)
	randomLifeText.update(screen)

	// Life numberInput
	life.disabled = (config.General.HaveLife && config.General.RandomLife) || !config.General.HaveLife
	life.update(screen)
	lifeText.update(screen)

	// screenManagement title
	screenManagement.update(screen)

	// marginOutsideScreen slider
	marginOutsideScreen.update(screen)
	marginOutsideScreenValue.update(screen)
	marginOutsideScreenText.update(screen)

	// dynamicChanges title
	dynamicChangesText.update(screen)
	// changeColorAccordingToTitle
	changeColorAccordingToTitle.update(screen)

	// changeColorAccordingToNothing
	changeColorAccordingToNothing.checked = config.General.ChangeColorAccordingTo == 0
	changeColorAccordingToNothing.update(screen)
	changeColorAccordingToNothingText.update(screen)

	// changeColorAccordingToPosition
	changeColorAccordingToPosition.checked = config.General.ChangeColorAccordingTo == 1
	changeColorAccordingToPosition.update(screen)
	changeColorAccordingToPositionText.update(screen)

	// changeColorAccordingToLife
	changeColorAccordingToLife.checked = config.General.ChangeColorAccordingTo == 2
	changeColorAccordingToLife.disabled = !config.General.HaveLife
	changeColorAccordingToLife.update(screen)
	if config.General.ChangeColorAccordingTo == 2 && !config.General.HaveLife {
		config.General.ChangeColorAccordingTo = 0
	}
	changeColorAccordingToLifeText.update(screen)

	// minColorRed
	minColorRed.update(screen)
	minColorRedValue.update(screen)
	minColorRedText.update(screen)

	// minColorGreen
	minColorGreen.update(screen)
	minColorGreenValue.update(screen)
	minColorGreenText.update(screen)

	// minColorBlue
	minColorBlue.update(screen)
	minColorBlueValue.update(screen)
	minColorBlueText.update(screen)

	// maxColorRed
	maxColorRed.update(screen)
	maxColorRedValue.update(screen)
	maxColorRedText.update(screen)

	// maxColorGreen
	maxColorGreen.update(screen)
	maxColorGreenValue.update(screen)
	maxColorGreenText.update(screen)

	// maxColorBlue
	maxColorBlue.update(screen)
	maxColorBlueValue.update(screen)
	maxColorBlueText.update(screen)

	// changeScaleAccordingToTitle
	changeScaleAccordingToTitle.update(screen)

	// changeScaleAccordingToNothing
	changeScaleAccordingToNothing.checked = config.General.ChangeScaleAccordingTo == 0
	changeScaleAccordingToNothing.update(screen)
	changeScaleAccordingToNothingText.update(screen)

	// changeScaleAccordingToPosition
	changeScaleAccordingToPosition.checked = config.General.ChangeScaleAccordingTo == 1
	changeScaleAccordingToPosition.update(screen)
	changeScaleAccordingToPositionText.update(screen)

	// changeScaleAccordingToLife
	changeScaleAccordingToLife.checked = config.General.ChangeScaleAccordingTo == 2
	changeScaleAccordingToLife.disabled = !config.General.HaveLife
	changeScaleAccordingToLife.update(screen)
	changeScaleAccordingToLifeText.update(screen)

	// changeOpacityAccordingToTitle
	changeOpacityAccordingToTitle.update(screen)

	// changeOpacityAccordingToNothing
	changeOpacityAccordingToNothing.checked = config.General.ChangeOpacityAccordingTo == 0
	changeOpacityAccordingToNothing.update(screen)
	changeOpacityAccordingToNothingText.update(screen)

	// changeOpacityAccordingToPosition
	changeOpacityAccordingToPosition.checked = config.General.ChangeOpacityAccordingTo == 1
	changeOpacityAccordingToPosition.update(screen)
	changeOpacityAccordingToPositionText.update(screen)

	// changeOpacityAccordingToLife
	changeOpacityAccordingToLife.checked = config.General.ChangeOpacityAccordingTo == 2
	changeOpacityAccordingToLife.disabled = !config.General.HaveLife
	changeOpacityAccordingToLife.update(screen)
	changeOpacityAccordingToLifeText.update(screen)

	// Met à jour l'état des bouton et l'affiche
	// Si le bouton playButton n'est plus à la bonne place en fonction de la taille d'écran
	if PlayButton.x < config.General.WindowSizeX-PlayButton.width || PlayButton.x > config.General.WindowSizeX-PlayButton.width+30 || PlayButton.y < config.General.WindowSizeY-PlayButton.height || PlayButton.y > config.General.WindowSizeY-PlayButton.height+30 {
		PlayButton.x = config.General.WindowSizeX - PlayButton.width - 30
		PlayButton.y = config.General.WindowSizeY - PlayButton.height - 130
	}
	PlayButton.update(screen)
	// Si le bouton accessParticlesButton n'est plus à la bonne place en fonction de la taille d'écran
	if accessParticlesButton.x < config.General.WindowSizeX-accessParticlesButton.width || accessParticlesButton.x > config.General.WindowSizeX-accessParticlesButton.width+30 || accessParticlesButton.y < config.General.WindowSizeY-accessParticlesButton.height || accessParticlesButton.y > config.General.WindowSizeY-accessParticlesButton.height+30 {
		accessParticlesButton.x = config.General.WindowSizeX - accessParticlesButton.width - 30
		accessParticlesButton.y = config.General.WindowSizeY - accessParticlesButton.height - 80
	}
	accessParticlesButton.update(screen)

	// Si le bouton leaveGamebutton n'est plus à la bonne place en fonction de la taille d'écran
	if leaveGamebutton.x < config.General.WindowSizeX-leaveGamebutton.width || leaveGamebutton.x > config.General.WindowSizeX-leaveGamebutton.width+30 || leaveGamebutton.y < config.General.WindowSizeY-leaveGamebutton.height || leaveGamebutton.y > config.General.WindowSizeY-leaveGamebutton.height+30 {
		leaveGamebutton.x = config.General.WindowSizeX - leaveGamebutton.width - 30
		leaveGamebutton.y = config.General.WindowSizeY - leaveGamebutton.height - 30
	}
	leaveGamebutton.update(screen)

	return nil
}
