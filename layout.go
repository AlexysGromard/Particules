package main

import (
	"project-particles/config"
)

// Layout définit la taille en pixels de la zone d'affichage des particules
// en fonction de la taille en pixels de la fenêtre. Vous n'avez jamais à
// modifier cette fonction.
func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.General.WindowSizeX, config.General.WindowSizeY
}
