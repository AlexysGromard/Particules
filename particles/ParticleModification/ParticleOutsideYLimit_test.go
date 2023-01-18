package ParticleModification

import (
	"project-particles/particles/Test"
	"testing"
)

func TestParticleOutsideY1(t *testing.T) {
	MaxPositionY := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionY = 600
	resultat := ParticleOutsideYLimit(&particule, MaxPositionY, MarginOutside)

	if !resultat {
		t.Error("la particule est en dehore dès limite mais la fonction du qu'elle est dans les limites")
	}
}

func TestParticleOutsideY2(t *testing.T) {
	MaxPositionY := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionY = -100
	resultat := ParticleOutsideYLimit(&particule, MaxPositionY, MarginOutside)

	if !resultat {
		t.Error("la particule est en dehore dès limite mais la fonction du qu'elle est dans les limites")
	}
}

func TestParticleinsideY1(t *testing.T) {
	MaxPositionY := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionY = 200
	resultat := ParticleOutsideYLimit(&particule, MaxPositionY, MarginOutside)

	if resultat {
		t.Error("la particule est dans les limites mais la fonction du qu'elle est en dehore des limites")
	}
}
