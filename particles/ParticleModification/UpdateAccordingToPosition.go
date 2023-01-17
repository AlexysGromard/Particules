package ParticleModification

import (
	"project-particles/particles"
)

//cette fonction renvois une valeur en fonction de la position

//entrer : ValeurMin la valeur minimal que peux retrouner la fonction
//		 : ValeurMax la veleur maximal que peux retrouner la fonction
func UpdateAccordingToPosition(p *particles.Particle, ValeuMin, ValeurMax float64, MaxPosition int) float64 {
	return ((p.PositionX+p.PositionY)/float64(MaxPosition))*(ValeurMax-ValeuMin) + ValeuMin
}
