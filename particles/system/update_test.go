package system

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestUpdateGravity vérifie que la gravité est bien appliqué
// Elle vérifie aussi que la vitesse en X n'est pas modifier
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

// Le test TestUpdateVitesse vérifie que la vitesse est bien appliqué
// Elle vérifie aussi que la position en X et Y a bien été modifier
// et que la vitesse a bien été modifié.
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
		t.Error("la vitesse n'a pas été appliqué")
	} else if particuleEnd.PositionX != 100 || particuleEnd.PositionY != 90 {
		t.Error("la vitesse n'a pas été correctement appliqué")
	}

}

func TestUpdateLifeFalse(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	Particule := Test.Basique_Particule()
	LifeStart := 50
	Particule.Life = LifeStart

	sys.Content.PushFront(&Particule)

	config.General.HaveLife = false

	sys.Update()

	particuleEnd := sys.Content.Front().Value.(*particles.Particle)

	if particuleEnd.Life != LifeStart {
		t.Error("La vie à été modifier alors que le parametre haveLife est à false")
	}
}

func TestUpdateLifeTrue(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	Particule := Test.Basique_Particule()
	LifeStart := 50
	Particule.Life = LifeStart

	sys.Content.PushFront(&Particule)

	config.General.HaveLife = true

	sys.Update()

	particuleEnd := sys.Content.Front().Value.(*particles.Particle)

	if particuleEnd.Life != LifeStart-1 {
		t.Error("La vie n'a pas correctement diminuer alors que le parametre HaveLife est à true")
	}
}

func TestUpdateKillParticuleIfOutside1(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	Particule := Test.Basique_Particule()
	Particule.PositionX, Particule.PositionY = 100, 100
	Particule.Life = 50

	sys.Content.PushFront(&Particule)

	config.General.WindowSizeX = 200
	config.General.WindowSizeY = 200
	config.General.MarginOutsideScreen = 1

	sys.Update()

	//particuleEnd := sys.Content.Front().Value.(*particles.Particle)

	if sys.Content.Len() != 1 {
		t.Error("La vie n'a pas correctement diminuer alors que le parametre HaveLife est à true")
	}
}
