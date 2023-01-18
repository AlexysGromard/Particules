package ParticleModification

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestMakeWhirlwind0 vérifie que la fonction MakeWhirlwind ne modifie pas la vitesse d'une particule si elle est à 0 de distance par raport au centre
func TestMakeWhirlwind0(t *testing.T) {
	CentreX, CentreY := 100, 100

	l := list.New()

	particule := Test.Basique_Particule()
	particule.PositionX, particule.PositionY = 100, 100

	SpeedInitX, SpeedInitY := float64(10), float64(10)
	particule.SpeedX, particule.SpeedY = SpeedInitX, SpeedInitY

	l.PushFront(&particule)

	MakeWhirlwind(l, 1, float64(CentreX), float64(CentreY))

	if l.Front().Value.(*particles.Particle).SpeedX != SpeedInitX || l.Front().Value.(*particles.Particle).SpeedY != SpeedInitY {
		t.Error("la distance par raport au centre est de 0 alors la vitesse ne devrais pas être modifier mais elle est différente de savitesse initial")
	}
}

// Le test TestMakeWhirlwindSpeedType0 vérifie que la fonction MakeWhirlwind ne modifie pas la vitesse d'une particule si le TypeSpeed est à 0
// Si ça n'est pas le cas, le test échoue
func TestMakeWhirlwindSpeedType0(t *testing.T) {
	CentreX, CentreY := 100, 100

	l := list.New()

	particule := Test.Basique_Particule()
	particule.PositionX, particule.PositionY = 20, 20

	SpeedInitX, SpeedInitY := float64(10), float64(10)
	particule.SpeedX, particule.SpeedY = SpeedInitX, SpeedInitY

	l.PushFront(&particule)

	MakeWhirlwind(l, 0, float64(CentreX), float64(CentreY))

	if l.Front().Value.(*particles.Particle).SpeedX != 0 || l.Front().Value.(*particles.Particle).SpeedY != 0 {
		t.Error("le TypeSpeed est à 0 alors les vitesse X et y devrait être à 0")
	}
}

// Le test TestMakeWhirlwindSpeedType1 vérifie que la fonction MakeWhirlwind modifie la vitesse d'une particule si le TypeSpeed est à 1
func TestMakeWhirlwindSpeedType1(t *testing.T) {
	CentreX, CentreY := 100, 100

	l := list.New()

	particule := Test.Basique_Particule()
	particule.PositionX, particule.PositionY = 20, 20

	SpeedInitX, SpeedInitY := float64(0), float64(0)
	particule.SpeedX, particule.SpeedY = SpeedInitX, SpeedInitY

	l.PushFront(&particule)

	MakeWhirlwind(l, 1, float64(CentreX), float64(CentreY))

	if l.Front().Value.(*particles.Particle).SpeedX == 0 || l.Front().Value.(*particles.Particle).SpeedY == 0 {
		t.Error("le TypeSpeed n'est pas 0 alors les vitesse X ou y devrait toujour être à 0")
	}
}
