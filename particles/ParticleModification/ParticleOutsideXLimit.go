package ParticleModification

import (
	"project-particles/particles"
)

//on considaire que la position minimal en X est de 0
func ParticleOutsideXLimit(p *particles.Particle, MaxPositionX,MarginOutsideScreen int) bool {
	return (p.PositionX < float64(0-MarginOutsideScreen) || p.PositionX > float64(MaxPositionX+MarginOutsideScreen))
}