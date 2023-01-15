package particles


//on considaire que la position minimal en X est de 0
func(p *Particle) ParticleOutsideXLimit(MaxPositionX,MarginOutsideScreen int) bool {
	return (p.PositionX < float64(0-MarginOutsideScreen) || p.PositionX > float64(MaxPositionX+MarginOutsideScreen))
}