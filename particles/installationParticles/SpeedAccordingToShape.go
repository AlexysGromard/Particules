package installationParticles

import (
	"math"
	"math/rand"
)

// La fonction SpeedAccordingToShape retourne la vitesse en X et en Y en fonction de la forme
// géométrique choisie
func SpeedAccordingToShape(SpeedMode int, positionX, positionY, centreX, centreY float64) (SpeedX, SpeedY float64) {
	if SpeedMode == 0 {
		return 0, 0
	}

	maxSpeed, minSpeed := minAndMaxSpeed(SpeedMode)

	// Récuépartion de l'angle en fonction du centre
	var angle float64 = math.Atan2(positionY-float64(centreY), positionX-float64(centreX))
	// calcule de la vitesse
	var vitesse float64 = (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64())
	// Calcul de la vitesse en fonction de l'angle
	SpeedX = vitesse * math.Cos(angle)
	SpeedY = vitesse * math.Sin(angle)

	return SpeedX, SpeedY
}
