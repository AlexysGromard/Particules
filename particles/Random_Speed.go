package particles

import (
	"math"
	"math/rand"

)

// La fonction Random_Speed permet de définir la vitesse de la particule aléatoirement
// Valeur d'entrée : un entier (mode de vitesse)
// Valeur de sortie : deux flottants (vitesse de la particule x et y)
func Random_Speed(SpeedMode int) (SpeedX, SpeedY float64) {

	if SpeedMode == 0 {
		return 0, 0
	}

	// Initialisation de l'intervalle vitesse de la particule
	maxSpeed, minSpeed := MinAndMaxSpeed(SpeedMode)
	// Initialisation de la direction de la particule
	var direction float64 = float64(rand.Intn(360))
	// Initialisation de la vitesse de la particule
	var vitesse float64 = float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()
	// Calcul de la vitesse de la particule
	SpeedX = math.Cos(direction) * vitesse
	SpeedY = math.Sin(direction) * vitesse

	return SpeedX, SpeedY
}
