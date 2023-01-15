package particles

import (
	"testing"
)

func Test1AccordingToPosition(t *testing.T) {
	var particule Particle
	particule.PositionX, particule.PositionY = 0, 0
	LargeurFenaitre, HauteurFenaitre := 100, 100
	var minCouleur, maxCouleur float64 = 0, 1

	particule = Basique_Particule()
	couleur_init := particule.UpdateAccordingToPosition(minCouleur, maxCouleur, LargeurFenaitre+HauteurFenaitre)

	particule.PositionX, particule.PositionY = 50, 30
	couleur_fin := particule.UpdateAccordingToPosition(minCouleur, maxCouleur, LargeurFenaitre+HauteurFenaitre)

	if couleur_init == couleur_fin {
		t.Error("la couleur n'a pas changer alors que la position à changer")
	}
}

func Test2AccordingToPosition(t *testing.T) {
	var particule Particle
	LargeurFenaitre, HauteurFenaitre := 100, 100
	var minCouleur, maxCouleur float64 = 0, 1

	particule = Basique_Particule()
	particule.PositionX, particule.PositionY = 50, 50
	particule.ColorRed, particule.ColorGreen ,particule.ColorBlue = 1, 1, 1

	couleur := particule.UpdateAccordingToPosition(minCouleur, maxCouleur, LargeurFenaitre+HauteurFenaitre)

	if couleur != 0.5 {
		t.Error("la couleur n'a pas changer alors que la position à changer")
	}
}
