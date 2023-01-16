package particles

import (
	"math/rand"

)

// La fonction Random_Position permet de définir la position de la particule aléatoirement
// Valeur d'entrée : deux entiers (largeur et hauteur)
// Valeur de sortie : deux flottants (coordonnées de la particule)
func Random_Position(intervalX, intervalY int) (PositionX, PositionY float64) {
	// On choisit un nombre aléatoire entre 0 et l'intervalle
	PositionX = float64(rand.Intn(intervalX)) + rand.Float64()
	PositionY = float64(rand.Intn(intervalY)) + rand.Float64()
	return PositionX, PositionY
}
