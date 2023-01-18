package ParticleModification

import (
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test Test1AccordingToPosition vérifie que la fonction UpdateAccordingToPosition renvoie une couleur différente si la particule a changé de position
// Si ce n'est pas le cas, le test échoue
func Test1AccordingToPosition(t *testing.T) {
	var particule particles.Particle
	particule.PositionX, particule.PositionY = 0, 0
	LargeurFenaitre, HauteurFenaitre := 100, 100
	var minCouleur, maxCouleur float64 = 0, 1

	particule = Test.Basique_Particule()
	couleur_init := UpdateAccordingToPosition(&particule, minCouleur, maxCouleur, LargeurFenaitre+HauteurFenaitre)

	particule.PositionX, particule.PositionY = 50, 30
	couleur_fin := UpdateAccordingToPosition(&particule, minCouleur, maxCouleur, LargeurFenaitre+HauteurFenaitre)

	if couleur_init == couleur_fin {
		t.Error("la couleur n'a pas changer alors que la position à changer")
	}
}

// Le test Test2AccordingToPosition vérifie que la fonction UpdateAccordingToPosition renvoie une couleur différente si la particule a changé de position
// Si ce n'est pas le cas, le test échoue
func Test2AccordingToPosition(t *testing.T) {
	var particule particles.Particle
	LargeurFenaitre, HauteurFenaitre := 100, 100
	var minCouleur, maxCouleur float64 = 0, 1

	particule = Test.Basique_Particule()
	particule.PositionX, particule.PositionY = 50, 50
	particule.ColorRed, particule.ColorGreen, particule.ColorBlue = 1, 1, 1

	couleur := UpdateAccordingToPosition(&particule, minCouleur, maxCouleur, LargeurFenaitre+HauteurFenaitre)

	if couleur != 0.5 {
		t.Error("la couleur n'a pas changer alors que la position à changer")
	}
}
