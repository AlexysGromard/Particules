package particles

import "testing"

func TestParticleOutsideY1(t *testing.T){
	MaxPositionY := 500
	MarginOutside := 20

	particule := Basique_Particule()
	particule.PositionY = 600
	res := particule.ParticleOutsideYLimit(MaxPositionY,MarginOutside)

	if !res{
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

func TestParticleOutsideY2(t *testing.T){
	MaxPositionY := 500
	MarginOutside := 20

	particule := Basique_Particule()
	particule.PositionY = -100
	res := particule.ParticleOutsideYLimit(MaxPositionY,MarginOutside)

	if !res{
		t.Error("la particule est en dehore des limite mais la fonction du qu'elle est dans les limite")
	}
}

func TestParticleinsideY1(t *testing.T){
	MaxPositionY := 500
	MarginOutside := 20

	particule := Basique_Particule()
	particule.PositionY = 200
	res := particule.ParticleOutsideYLimit(MaxPositionY,MarginOutside)

	if res{
		t.Error("la particule est dans les limite mais la fonction du qu'elle est en dehore des limite")
	}
}