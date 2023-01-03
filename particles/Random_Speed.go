package particles

import (
	"math/rand"
)

func Random_Speed(SpeedMode int) (SpeedX, SpeedY float64) {
	// Initialisation de la vitesse de la particule
	maxSpeed, minSpeed := MinAndMaxSpeed(SpeedMode)
	// Initialisation de la vitesse de la particule
	SpeedX = (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1)
	SpeedY = (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1)

	return SpeedX, SpeedY
}
