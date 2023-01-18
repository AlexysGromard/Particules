package ParticleModification

import (
	"project-particles/particles"
)

//cette fonction renvois une valeur en fonction de la Vie de la patucle

// Valeurs d'éntrées : ValeurMin (la valeur minimal que peux retrouner la fonction),
// ValeurMax (la veleur maximal que peux retrouner la fonction)
func UpdateAccordingToLife(p *particles.Particle, ValeurMin, ValeurMax float64) float64 {
	return (float64(p.Life)/float64(p.LifeInit))*(ValeurMax-ValeurMin) + ValeurMin
}
