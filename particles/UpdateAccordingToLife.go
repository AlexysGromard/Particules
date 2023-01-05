package particles

func (p *Particle) UpdateAccordingToLife(ValeuMin, ValeurMax float64) float64 {
	return (float64(p.Life)/float64(p.LifeInit))*(ValeurMax-ValeuMin) + ValeuMin
}
