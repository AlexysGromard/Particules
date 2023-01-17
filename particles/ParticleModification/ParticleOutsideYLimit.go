package ParticleModification

import (
	"project-particles/particles"
)

//on considaire que la position minimal en Y est de 0
func ParticleOutsideYLimit(p *particles.Particle, MaYPositionY,MarginOutsideScreen int) bool {
	return (p.PositionY < float64(0-MarginOutsideScreen) || p.PositionY > float64(MaYPositionY+MarginOutsideScreen))
}