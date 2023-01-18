package main

import (
	"image"
	_ "image/png"
	"log"
	"os"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/configPage"
	"project-particles/particles/system"

	"github.com/hajimehoshi/ebiten/v2"
)

// main est la fonction principale du projet. Elle commence par lire le fichier
// de configuration, puis elle charge en mémoire l'image d'une particule. Elle
// initialise ensuite la fenêtre d'affichage, puis elle crée un système de
// particules encapsulé dans un "game" et appelle la fonction RunGame qui se
// charge de faire les mise-à-jour (Update) et affichages (Draw) de manière
// régulière.
func main() {
	config.Get("config.json")
	assets.Get()

	ebiten.SetWindowTitle(config.General.WindowTitle)
	ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)
	_, _, maxW, maxH := ebiten.WindowSizeLimits()                  // Taille minimum et max de la fentre
	ebiten.SetWindowSizeLimits(1024, 500, maxW, maxH)              // Taille minimum et max de la fentre
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled) // Autorise le redimensionnement de la fentre
	// Ajouter une icône à la fenêtre
	// Charger l'icon ./assets/icon.png
	//load
	file16, _ := os.Open("./assets/icon_16.png")
	file32, _ := os.Open("./assets/icon_32.png")
	file64, _ := os.Open("./assets/icon_64.png")

	// Convert in image.Image
	img16, _, _ := image.Decode(file16)
	img32, _, _ := image.Decode(file32)
	img64, _, _ := image.Decode(file64)

	// Convert in ebiten.Image
	icon16 := ebiten.NewImageFromImage(img16)
	icon32 := ebiten.NewImageFromImage(img32)
	icon64 := ebiten.NewImageFromImage(img64)
	// image.Image
	icon := []image.Image{icon16, icon32, icon64}
	ebiten.SetWindowIcon(icon)

	// Chargement des images de configuration
	configPage.LoadAllImages("./assets/graphics/")

	g := game{system: system.NewSystem()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
