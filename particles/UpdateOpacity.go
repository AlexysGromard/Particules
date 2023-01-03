package particles

import (
	"project-particles/config"
	//"fmt"
)

// Cette fonction actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeOpacityAccordingTo" dans config.json
func (p *Particle) UpdateOpacity(Mode int) {
	if config.General.ChangeOpacityAccordingTo == 1 {
		p.Opacity = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.Opacity
	} else if config.General.ChangeOpacityAccordingTo == 2 {
		p.Opacity = (float64(p.Life) / float64(config.General.Life)) * config.General.Opacity
	}
}
