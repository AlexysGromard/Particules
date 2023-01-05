package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
)

// Explosion
func abs(nb float64) float64 {
	if nb < 0 {
		return -nb
	}
	return nb
}

func speed() (speedX, speedY float64) {
	speedX = (float64(rand.Intn(5)) + rand.Float64()) * float64(rand.Intn(2)*2-1)
	speedY = (float64(rand.Intn(5)) + rand.Float64()) * float64(rand.Intn(2)*2-1)
	if abs(speedX)+abs(speedY) <= 2.5 {
		speedX, speedY = speed()
	}
	return speedX, speedY
}

func Explosion(l *list.List) {
	// On crée une explosion de particules
	for i := 0; i < 100; i++ {
		SpeedX, SpeedY := speed()
		l.PushFront(&Particle{
			PositionX: float64(config.General.SpawnX),
			PositionY: float64(config.General.SpawnY),
			Rotation:  config.General.Rotation,
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
			ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
			Opacity: config.General.Opacity,
			SpeedX:  SpeedX,
			SpeedY:  SpeedY,
			Life:    config.General.Life,
		})
	}
}
