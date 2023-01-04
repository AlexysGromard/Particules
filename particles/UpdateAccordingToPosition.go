package particles

func (p *Particle) UpdateAccordingToPosition(ValeuMin, ValeurMax float64, MaxPosition int) float64 {
	return ((p.PositionX+p.PositionY)/float64(MaxPosition))*(ValeurMax-ValeuMin) + ValeuMin
}
