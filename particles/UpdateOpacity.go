package particles

import (
	"project-particles/config"
	//"fmt"
)

// Cette fonction actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeOpacityAccordingTo" dans config.json
// Valeur d'entrée : un entier (mode de couleur)
// Valeur de sortie : aucune
// Cette fonction interagit directement avec la structure Particle grâce au pointeur
func (p *Particle) UpdateOpacity(mode int) {
	if mode == 1 { // mode 1 : couleur définie en fonction de la position de la particule
		p.Opacity = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.Opacity
	} else if mode == 2 { // mode 2 : couleur définie en fonction de la durée de vie de la particule
		p.Opacity = (float64(p.Life) / float64(config.General.Life)) * config.General.Opacity
	}
}
