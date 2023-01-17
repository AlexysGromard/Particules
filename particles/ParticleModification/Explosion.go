package ParticleModification

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"project-particles/particles/InstallationParticles"
)

func Explosion(l *list.List) {
	// On crée une explosion de particules
	for i := 0; i < 100; i++ {
		SpeedX, SpeedY := InstallationParticles.Random_Speed(3)
		l.PushFront(&particles.Particle{
			PositionX: float64(config.General.SpawnX),
			PositionY: float64(config.General.SpawnY),
			Rotation:  config.General.Rotation,
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
			ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
			Opacity:  config.General.Opacity,
			SpeedX:   SpeedX,
			SpeedY:   SpeedY,
			LifeInit: config.General.Life, Life: config.General.Life,
		})
	}
}