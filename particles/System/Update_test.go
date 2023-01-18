package System

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

func TestUpdateGravity(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	ParticuleInit := Test.Basique_Particule()
	SpeedXInit, SpeedYInit := float64(20), float64(20)
	ParticuleInit.SpeedX, ParticuleInit.SpeedY = SpeedXInit, SpeedYInit

	sys.Content.PushFront(&ParticuleInit)

	config.General.Gravity = 1

	sys.Update()

	if sys.Content.Front().Value.(*particles.Particle).SpeedX != SpeedXInit {
		t.Error("la vitesse en X à été modifier mais elle doit reter la même")
	} else if sys.Content.Front().Value.(*particles.Particle).SpeedY == SpeedYInit {
		t.Error("")
	} else if sys.Content.Front().Value.(*particles.Particle).SpeedY != SpeedYInit+config.General.Gravity {
		t.Error("la Gravite n'a pas été appliqué correctement")
	}
}

func TestUpdateVitesse(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	ParticuleInit := Test.Basique_Particule()
	ParticuleInit.SpeedX, ParticuleInit.SpeedY = 20, 10
	ParticuleInit.PositionX, ParticuleInit.PositionY = 80, 80

	sys.Content.PushFront(&ParticuleInit)

	config.General.Gravity = 0

	sys.Update()

	particuleEnd := sys.Content.Front().Value.(*particles.Particle)

	if particuleEnd.PositionX == ParticuleInit.PositionX || particuleEnd.PositionY == ParticuleInit.PositionY {
		t.Error("la vitesse n'a pas été appliqué")
	} else if particuleEnd.PositionX != 100 || particuleEnd.PositionY != 90 {
		t.Error("la vitesse n'a pas été correctement appliqué")
	}
}
