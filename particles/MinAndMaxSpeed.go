package particles

// MinAndMaxSpeed retourne la vitesse minimale et maximale en fonction du mode de vitesse
// choisi par l'utilisateur.
// Valeur d'entree : un entier compris entre 0 et 3
// Valeur de sortie : deux entiers compris entre 0 et 5 qui correspondent a la vitesse minimale
// et maximale
func MinAndMaxSpeed(SpeedMode int) (minSpeed, maxSpeed int) {
	switch SpeedMode {
	// -1: vitesse aleatoire entre 0 et 5
	case -1:
		minSpeed = 0
		maxSpeed = 5
	// 1: vitesse aleatoire entre 0 et 1
	case 1:
		minSpeed = 0
		maxSpeed = 1
	// 2: vitesse aleatoire entre 1 et 3
	case 2:
		minSpeed = 1
		maxSpeed = 3
	// 3: vitesse aleatoire entre 3 et 5
	case 3:
		minSpeed = 3
		maxSpeed = 5
	}

	return maxSpeed, minSpeed
}
