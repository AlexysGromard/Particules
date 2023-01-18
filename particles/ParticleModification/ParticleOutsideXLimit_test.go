package ParticleModification

import (
	"project-particles/particles/Test"
	"testing"
)

// Le test TestParticleOutsideX1 vérifie que la fonction ParticleOutsideXLimit renvoie true si la particule est en dehors des limites
// Si ce n'est pas le cas, le test échoue
func TestParticleOutsideX1(t *testing.T) {
	MaxPositionX := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionX = 600
	res := ParticleOutsideXLimit(&particule, MaxPositionX, MarginOutside)

	if !res {
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

// Le test TestParticleOutsideX2 vérifie que la fonction ParticleOutsideXLimit renvoie true si la particule est en dehors des limites
// Si ça n'est pas le cas, le test échoue
func TestParticleOutsideX2(t *testing.T) {
	MaxPositionX := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionX = -100
	res := ParticleOutsideXLimit(&particule, MaxPositionX, MarginOutside)

	if !res {
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

// Le test TestParticleinsideX1 vérifie que la fonction ParticleOutsideXLimit renvoie false si la particule est dans les limites
// Si ce n'est pas le cas, le test échoue
func TestParticleinsideX1(t *testing.T) {
	MaxPositionX := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionX = 200
	res := ParticleOutsideXLimit(&particule, MaxPositionX, MarginOutside)

	if res {
		t.Error("la particule est dans les limite mais la fonction du qu'elle est en dehore des limite")
	}
}
