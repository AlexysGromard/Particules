package particles

import (
	"project-particles/config"
)

// Cette fonction actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeRotationAccordingTo" dans config.json
func (p *Particle) UpdateRotation(Mode int) {
	if Mode == 1 {
		p.Rotation = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.Rotation
	} else if Mode == 2 {
		p.Rotation = (float64(p.Life) / float64(config.General.Life)) * config.General.Rotation
	}
}
