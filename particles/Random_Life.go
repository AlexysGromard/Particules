package particles

import (
	"math/rand"

)

// La fonction Random_Life permet de définir la durée de vie de la particule aléatoirement
// Valeur d'entrée : lifeMax un entier positif ou nul qui défini la vie maximume que peux générer la fonction 
// Valeur de sortie : un entier compris entre 0 et 50
func Random_Life(lifeMax uint) (life int) {
	life = rand.Intn(int(lifeMax))
	return life
}
