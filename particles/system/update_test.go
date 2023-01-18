package system

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

func TestUpdateGravity(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	Particule := Test.Basique_Particule()
	SpeedXInit, SpeedYInit := float64(20), float64(20)
	Particule.PositionX, Particule.PositionY = 100, 100
	Particule.SpeedX, Particule.SpeedY = SpeedXInit, SpeedYInit

	sys.Content.PushFront(&Particule)

	config.General.Gravity = 1
	config.General.WindowSizeX, config.General.WindowSizeY = 1000, 1000

	sys.Update()

	particuleEnd, _ := sys.Content.Front().Value.(*particles.Particle)

	if particuleEnd.SpeedX != SpeedXInit {
		t.Error("la vitesse en X à été modifier mais elle doit reter la même")
	} else if particuleEnd.SpeedY == SpeedYInit {
		t.Error("")
	} else if particuleEnd.SpeedY != SpeedYInit+config.General.Gravity {
		t.Error("la Gravite n'a pas été appliqué correctement")
	}
}

func TestUpdateVitesse(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	Particule := Test.Basique_Particule()
	Particule.SpeedX, Particule.SpeedY = 20, 10
	PositionXInit, PositionYInit := float64(80), float64(80)
	Particule.PositionX, Particule.PositionY = PositionXInit, PositionYInit

	sys.Content.PushFront(&Particule)

	config.General.Gravity = 0
	config.General.WindowSizeX, config.General.WindowSizeY = 1000, 1000

	sys.Update()

	particuleEnd := sys.Content.Front().Value.(*particles.Particle)

	if particuleEnd.PositionX == PositionXInit || particuleEnd.PositionY == PositionYInit {
		t.Error("la vitesse n'a pas été appliqué", Particule.PositionX, Particule.PositionY)
	} else if particuleEnd.PositionX != 100 || particuleEnd.PositionY != 90 {
		t.Error("la vitesse n'a pas été correctement appliqué")
	}
}
