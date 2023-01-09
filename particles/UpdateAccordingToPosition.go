package particles


//cette fonction renvois une valeur en fonction de la position

//entrer : ValeurMin la valeur minimal que peux retrouner la fonction
//		 : ValeurMax la veleur maximal que peux retrouner la fonction
func (p *Particle) UpdateAccordingToPosition(ValeuMin, ValeurMax float64, MaxPosition int) float64 {
	return ((p.PositionX+p.PositionY)/float64(MaxPosition))*(ValeurMax-ValeuMin) + ValeuMin
}
