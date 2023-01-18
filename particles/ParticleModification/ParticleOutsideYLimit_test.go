package ParticleModification

import (
	"project-particles/particles/Test"
	"testing"
)

// Le test TestParticleOutsideY1 vérifie que la fonction ParticleOutsideYLimit renvoie true si la particule est en dehors des limites
// Si ce n'est pas le cas, le test échoue
func TestParticleOutsideY1(t *testing.T) {
	MaxPositionY := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionY = 600
	resultat := ParticleOutsideYLimit(&particule, MaxPositionY, MarginOutside)

	if !resultat {
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

// Le test TestParticleOutsideY2 vérifie que la fonction ParticleOutsideYLimit renvoie true si la particule est en dehors des limites
// Si ça n'est pas le cas, le test échoue
func TestParticleOutsideY2(t *testing.T) {
	MaxPositionY := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionY = -100
	resultat := ParticleOutsideYLimit(&particule, MaxPositionY, MarginOutside)

	if !resultat {
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

// Le test TestParticleinsideY1 vérifie que la fonction ParticleOutsideYLimit renvoie false si la particule est dans les limites
// Si ce n'est pas le cas, le test échoue
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
