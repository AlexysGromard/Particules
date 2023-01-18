package system

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

func TestCollisionWithWallFalse(t *testing.T) {

	l := list.New()

	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan

	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 100, 100
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1

	l.PushFront(&Particule1)

	CollisionWithWall(&Particule1, LimiteX, LimiteY)

	particuleEnd1, _ := l.Front().Value.(*particles.Particle)

	if particuleEnd1.SpeedX != SpeedXInit1 || particuleEnd1.SpeedY != SpeedYInit1 {
		t.Error("la vitesse de la particule à été modifier alors qu'elle n'a pas toucher les bords")
	}
}

func TestCollisionWithWallTrueUP(t *testing.T) {

	l := list.New()

	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan

	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 100, 0
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1

	l.PushFront(&Particule1)

	CollisionWithWall(&Particule1, LimiteX, LimiteY)

	particuleEnd1, _ := l.Front().Value.(*particles.Particle)

	if particuleEnd1.SpeedX != SpeedXInit1 || particuleEnd1.SpeedY != -SpeedYInit1 || particuleEnd1.PositionX != 100 || particuleEnd1.PositionY != 0.1 {
		t.Error("la vitesse de la particule à été modifier alors qu'elle n'a pas toucher les bords ou la particule n'a été éloigner de la bordure")
	}
}

func TestCollisionWithWallTrueLOW(t *testing.T) {

	l := list.New()

	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan

	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 100, 200
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1

	l.PushFront(&Particule1)

	CollisionWithWall(&Particule1, LimiteX, LimiteY)

	particuleEnd1, _ := l.Front().Value.(*particles.Particle)

	if particuleEnd1.SpeedX != SpeedXInit1 || particuleEnd1.SpeedY != -SpeedYInit1 || particuleEnd1.PositionX != 100 || particuleEnd1.PositionY != 199.9 {
		t.Error("la vitesse de la particule à été modifier alors qu'elle n'a pas toucher les bords ou la particule n'a été éloigner de la bordure")
	}
}

func TestCollisionWithWallTrueLEFT(t *testing.T) {

	l := list.New()

	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan

	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 0, 100
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1

	l.PushFront(&Particule1)

	CollisionWithWall(&Particule1, LimiteX, LimiteY)

	particuleEnd1, _ := l.Front().Value.(*particles.Particle)

	if particuleEnd1.SpeedX != -SpeedXInit1 || particuleEnd1.SpeedY != SpeedYInit1 || particuleEnd1.PositionX != 0.1 || particuleEnd1.PositionY != 100 {
		t.Error("la vitesse de la particule à été modifier alors qu'elle n'a pas toucher les bords ou la particule n'a été éloigner de la bordure")
	}
}

func TestCollisionWithWallTrueRIGHT(t *testing.T) {

	l := list.New()

	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan

	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 200, 100
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1

	l.PushFront(&Particule1)

	CollisionWithWall(&Particule1, LimiteX, LimiteY)

	particuleEnd1, _ := l.Front().Value.(*particles.Particle)

	if particuleEnd1.SpeedX != -SpeedXInit1 || particuleEnd1.SpeedY != SpeedYInit1 || particuleEnd1.PositionX != 199.9 || particuleEnd1.PositionY != 100 {
		t.Error("la vitesse de la particule à été modifier alors qu'elle n'a pas toucher les bords ou la particule n'a été éloigner de la bordure")
	}
}
