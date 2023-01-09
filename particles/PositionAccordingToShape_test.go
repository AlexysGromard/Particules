package particles

import (
	"math"
	"testing"
)

func TestPossitionAccordingToCircle(t *testing.T) {
	var centreX, centreY float64 = 100, 100
	largeur := 10
	posX, posY := PositionAccordingToShape("circle", largeur, centreX, centreY)

	DistanceO := math.Sqrt(carrée(posX-float64(centreX)) + carrée(posY-float64(centreY))) //La distence entre les cordonner générer par rapore au centre ici (100;100)
	//comme le calcule de DistanceO il y a des arondie le résultat n'est pas exactement le bon.
	//je déside donc d'arrondire le résultat à 5 chiffre après la virgule pour ne pas prendre en compte les arrondies.
	DistanceO = float64(int(DistanceO*100000) / 100000)

	if DistanceO != float64(largeur)/2 {
		t.Error("la distence par rapore à l'origine devrais être de ", largeur/2, " mais la distance est de ", DistanceO)
	}
}

func TestPossitionAccordingToSquare(t *testing.T) {
	var centreX, centreY float64 = 100, 100
	largeur := 20
	posX, posY := PositionAccordingToShape("square", largeur, centreX, centreY)

	DistanceOX := posX - centreX
	DistanceOY := posY - centreY

	if DistanceOX != float64(largeur)/2 && DistanceOY != float64(largeur)/2 {
		t.Error("l'une des cordonner x ou y devrait être égale à ", largeur/2, " mais elle sont égale à PositionX = ", DistanceOX, " PossitionY = ", DistanceOX)
	}
}
