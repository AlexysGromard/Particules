package InstallationParticles

import (
	"math"
	"math/rand"
)

func SpeedAccordingToShape(SpeedMode int, positionX, positionY, centreX, centreY float64) (SpeedX, SpeedY float64) {
	maxSpeed, minSpeed := MinAndMaxSpeed(SpeedMode)

	// Récuépartion de l'angle en fonction du centre
	var angle float64 = math.Atan2(positionY-float64(centreY), positionX-float64(centreX))
	// calcule de la vitesse
	var vitesse float64 = (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64())
	// Calcul de la vitesse en fonction de l'angle
	SpeedX = vitesse * math.Cos(angle)//(float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * math.Cos(angle)
	SpeedY = vitesse * math.Sin(angle)//(float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * math.Sin(angle)

	return SpeedX, SpeedY
}