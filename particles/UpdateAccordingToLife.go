package particles


//cette fonction renvois une valeur en fonction de la Vie de la patucle

//entrer : ValeurMin la valeur minimal que peux retrouner la fonction
//		 : ValeurMax la veleur maximal que peux retrouner la fonction
func (p *Particle) UpdateAccordingToLife(ValeurMin, ValeurMax float64) float64 {
	return (float64(p.Life)/float64(p.LifeInit))*(ValeurMax-ValeurMin) + ValeurMin
}
