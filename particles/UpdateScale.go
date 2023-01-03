package particles

import (
	"project-particles/config"
)

// Cette fonction actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeScaleAccordingTo" dans config.json
func (p *Particle) UpdateScale(Mode int) {
	if Mode == 1 {
		p.ScaleX = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ScaleX
		p.ScaleY = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ScaleY
	} else if Mode == 2 {
		p.ScaleX = (float64(p.Life) / float64(config.General.Life)) * config.General.ScaleX
		p.ScaleY = (float64(p.Life) / float64(config.General.Life)) * config.General.ScaleY
	}
}
