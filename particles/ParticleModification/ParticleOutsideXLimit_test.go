package ParticleModification

import (
	"project-particles/particles/Test"
	"testing"
)

func TestParticleOutsideX1(t *testing.T) {
	MaxPositionX := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionX = 600
	resultat := ParticleOutsideXLimit(&particule, MaxPositionX, MarginOutside)

	if !resultat {
		t.Error("la particule est en dehore dès limite mais la fonction du qu'elle est dans les limites")
	}
}

func TestParticleOutsideX2(t *testing.T) {
	MaxPositionX := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionX = -100
	resultat := ParticleOutsideXLimit(&particule, MaxPositionX, MarginOutside)

	if !resultat {
		t.Error("la particule est en dehore dès limite mais la fonction du qu'elle est dans les limites")
	}
}

func TestParticleinsideX1(t *testing.T) {
	MaxPositionX := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionX = 200
	resultat := ParticleOutsideXLimit(&particule, MaxPositionX, MarginOutside)

	if resultat {
		t.Error("la particule est dans les limites mais la fonction dit qu'elle est en dehore des limites")
	}
}
