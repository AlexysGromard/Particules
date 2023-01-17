package ParticleModification

import (
	"testing"
	"project-particles/particles/Test"
)


func TestParticleOutsideY1(t *testing.T){
	MaxPositionY := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionY = 600
	resultat := ParticleOutsideYLimit(&particule,MaxPositionY,MarginOutside)

	if !resultat{
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

func TestParticleOutsideY2(t *testing.T){
	MaxPositionY := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionY = -100
	resultat := ParticleOutsideYLimit(&particule,MaxPositionY,MarginOutside)

	if !resultat{
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

func TestParticleinsideY1(t *testing.T){
	MaxPositionY := 500
	MarginOutside := 20

	particule := Test.Basique_Particule()
	particule.PositionY = 200
	resultat := ParticleOutsideYLimit(&particule,MaxPositionY,MarginOutside)

	if resultat{
		t.Error("la particule est dans les limite mais la fonction du qu'elle est en dehore des limite")
	}
}