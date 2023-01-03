package particles

import (
	"math"
	"math/rand"
)

// La fonction Random_Speed permet de définir la vitesse de la particule aléatoirement
// Valeur d'entrée : un entier (mode de vitesse)
// Valeur de sortie : deux flottants (vitesse de la particule x et y)
func Random_Speed(SpeedMode int) (SpeedX, SpeedY float64) {
	// Initialisation de l'interval vitesse de la particule
	maxSpeed, minSpeed := MinAndMaxSpeed(SpeedMode)

	var direction float64 = float64(rand.Intn(360))

	var vitesse float64 = float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()
	// Initialisation de la vitesse de la particule
	SpeedX = math.Cos(direction) * vitesse
	SpeedY = math.Sin(direction) * vitesse

	return SpeedX, SpeedY
}
