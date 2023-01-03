package particles

import (
	"project-particles/config"
)

// Cette méthode actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeColorAccordingTo" dans config.json
func (p *Particle) UpdateColor(Mode int) {
	if Mode == 1 {
		p.ColorRed = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorRed
		p.ColorGreen = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorGreen
		p.ColorBlue = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorBlue
	} else if Mode == 2 {
		p.ColorRed = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorRed
		p.ColorGreen = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorGreen
		p.ColorBlue = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorBlue
	}
}
