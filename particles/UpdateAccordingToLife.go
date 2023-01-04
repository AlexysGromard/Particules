package particles

func (p *Particle) UpdateAccordingToLife(ValeuMin, ValeurMax float64, MaxLife int) float64 {
	return (float64(p.Life)/float64(MaxLife))*(ValeurMax-ValeuMin) + ValeuMin
}
