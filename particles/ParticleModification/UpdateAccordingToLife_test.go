package ParticleModification

import (
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

func Test1AccordingToLife(t *testing.T) {
	var particule particles.Particle
	var minCouleur, maxCouleur float64 = 0, 1

	particule = Test.Basique_Particule()
	couleur_init := UpdateAccordingToLife(&particule, minCouleur, maxCouleur)

	particule.Life = 10
	couleur_fin := UpdateAccordingToLife(&particule, minCouleur, maxCouleur)

	if couleur_init == couleur_fin {
		t.Error("la couleur n'a pas changé alors que la position à changer")
	}
}

func Test2AccordingToLife(t *testing.T) {
	var particule particles.Particle
	var minCouleur, maxCouleur float64 = 0, 1

	particule = Test.Basique_Particule()
	particule.Life = 25
	couleur := UpdateAccordingToLife(&particule, minCouleur, maxCouleur)

	if couleur != 0.5 {
		t.Error("la couleur n'a pas changé alors que la position à changer")
	}
}
