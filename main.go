package main

import (
	"log"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/configPage"
	"project-particles/particles"

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
	_, _, maxW, maxH := ebiten.WindowSizeLimits()
	ebiten.SetWindowSizeLimits(1024, 750, maxW, maxH)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Taille minimum de la fentre

	// Chargement des images de configuration
	configPage.LoadImages()

	g := game{system: particles.NewSystem()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
