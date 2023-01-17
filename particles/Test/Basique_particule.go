package Test

import "project-particles/particles"

func Basique_Particule() particles.Particle {
	return particles.Particle{
		PositionX: 0,
		PositionY: 0,
		Rotation:  0,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:  1,
		SpeedX:   1,
		SpeedY:   1,
		LifeInit: 50,
		Life:     50,
	}
}