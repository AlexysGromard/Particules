package installationParticles

import (
	"math"
	"testing"
)

// Le test TestPossitionAccordingToCircle vérifie que la fonction PositionAccordingToShape retourne bien des coordonnées
// qui sont dans un cercle de rayon largeur/2
// Si ce n'est pas le cas, le test échoue
func TestPossitionAccordingToCircle(t *testing.T) {
	var centreX, centreY float64 = 100, 100
	largeur := 10
	posX, posY := PositionAccordingToShape("circle", largeur, centreX, centreY)

	DistanceO := math.Sqrt(carrée(posX-float64(centreX)) + carrée(posY-float64(centreY))) //La distence entre les cordonner générer par rapore au centre ici (100;100)
	//comme le calcule de DistanceO il y a des arondie le résultat n'est pas exactement le bon.
	//je déside donc d'arrondire le résultat à 5 chiffre après la virgule pour ne pas prendre en compte les arrondies.
	DistanceO = float64(int(DistanceO*100000) / 100000)

	if DistanceO != float64(largeur)/2 {
		t.Error("la distance par  rapport à l'origine devrait être de ", largeur/2, " mais la distance est de ", DistanceO)
	}
}

// Le test TestPossitionAccordingToSquare vérifie que la fonction PositionAccordingToShape retourne bien des coordonnées
// qui sont dans un carré de largeur largeur
// Si ce n'est pas le cas, le test échoue
func TestPossitionAccordingToSquare(t *testing.T) {
	var centreX, centreY float64 = 100, 100
	largeur := 20
	posX, posY := PositionAccordingToShape("square", largeur, centreX, centreY)

	DistanceOX := math.Abs(posX - centreX)
	DistanceOY := math.Abs(posY - centreY)

	if DistanceOX != float64(largeur)/2 && DistanceOY != float64(largeur)/2 {
		t.Error("l'une des cordonné x où y devrait être égale à ", largeur/2, " mais elles sont égales à PositionX = ", DistanceOX, " PositionY = ", DistanceOY)
	}
}
