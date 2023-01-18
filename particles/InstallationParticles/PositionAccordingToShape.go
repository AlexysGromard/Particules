package installationParticles

import (
	"math"
	"math/rand"
)

// La fonction PositionAccordingToShape permet de définir la position de la particule
// en fonction de la forme choisie par l'utilisateur.
// Valeur d'entrée : une chaîne de caractères (type de forme), un entier (taille de cette forme),
// deux flottants (coordonnées du centre de la forme)
func PositionAccordingToShape(TypeShape string, Width int, CentreX, CentreY float64) (PositionX, PositionY float64) {
	// Initialisation de la position de la particule sur un objet géométrique
	switch TypeShape {
	// Création d'un cercle
	case "circle":
		// On  Choisit un angle aléatoire et on calcule les coordonnées de la particule
		var randomPosition int = rand.Intn(360)
		PositionX = CentreX + float64(float64(Width/2)*math.Cos(float64(randomPosition)))
		PositionY = CentreY + float64(float64(Width/2)*math.Sin(float64(randomPosition)))
	// Création d'un carré
	case "square":
		// On choisit une une face du carré et une position aléatoire sur cette face
		var randomAxe int = rand.Intn(2)
		switch randomAxe {
		case 0:
			PositionX = CentreX - float64(Width/2) + float64(rand.Intn(2)*Width)
			PositionY = CentreY - float64(Width/2) + float64(rand.Float64()*float64(Width))
		case 1:
			PositionX = CentreX - float64(Width/2) + float64(rand.Float64()*float64(Width))
			PositionY = CentreY - float64(Width/2) + float64(rand.Intn(2)*Width)
		}
	}
	return PositionX, PositionY
}
