package particles

import (
	"project-particles/config"
)

// La fonction UpdateColor permet de mettre à jour la couleur de la particule en fonction
// du mode choisi.
// Valeur d'entrée : un entier (mode de couleur)
// Valeur de sortie : aucune
// Cette fonction interagit directement avec la structure Particle grâce au pointeur
func (p *Particle) UpdateColor(Mode int) {
	if Mode == 1 { // Mode 1 : couleur définie en fonction de la position de la particule
		p.ColorRed = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorRed
		p.ColorGreen = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorGreen
		p.ColorBlue = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorBlue
	} else if Mode == 2 { // Mode 2 : couleur définie en fonction de la durée de vie de la particule
		p.ColorRed = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorRed
		p.ColorGreen = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorGreen
		p.ColorBlue = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorBlue
	}
}
