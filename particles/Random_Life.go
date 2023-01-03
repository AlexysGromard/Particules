package particles

import (
	//"container/list"

	"math/rand"
	//"fmt"
)

// La fonction Random_Life permet de définir la durée de vie de la particule aléatoirement
// Valeur d'entrée : acune
// Valeur de sortie : un entier compris entre 0 et 50
func Random_Life() (life int) {
	life = rand.Intn(50)
	return life
}
