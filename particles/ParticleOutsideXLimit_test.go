package particles

import "testing"

func TestParticleOutsideX1(t *testing.T){
	MaxPositionX := 500
	MarginOutside := 20

	particule := Basique_Particule()
	particule.PositionX = 600
	res := particule.ParticleOutsideXLimit(MaxPositionX,MarginOutside)

	if !res{
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

func TestParticleOutsideX2(t *testing.T){
	MaxPositionX := 500
	MarginOutside := 20

	particule := Basique_Particule()
	particule.PositionX = -100
	res := particule.ParticleOutsideXLimit(MaxPositionX,MarginOutside)

	if !res{
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

func TestParticleinsideX1(t *testing.T){
	MaxPositionX := 500
	MarginOutside := 20

	particule := Basique_Particule()
	particule.PositionX = 200
	res := particule.ParticleOutsideXLimit(MaxPositionX,MarginOutside)

	if res{
		t.Error("la particule est dans les limite mais la fonction du qu'elle est en dehore des limite")
	}
}