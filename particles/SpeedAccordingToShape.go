package particles

import (
	"math"
	"math/rand"
)

func SpeedAccordingToShape(SpeedMode int, positionX, positionY, centreX, centreY float64) (SpeedX, SpeedY float64) {
	maxSpeed, minSpeed := MinAndMaxSpeed(SpeedMode)

	// Récuépartion de l'angle en fonction du centre
	var angle float64 = math.Atan2(positionY-float64(centreY), positionX-float64(centreX))
	// Calcul de la vitesse en fonction de l'angle
	SpeedX = (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * math.Cos(angle)
	SpeedY = (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * math.Sin(angle)

	return SpeedX, SpeedY
}
