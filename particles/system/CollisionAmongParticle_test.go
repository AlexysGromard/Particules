package system

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

func TestCollisionAmongParticleFalse(t *testing.T) {

	l := list.New()

	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	PositionXInit1, PositionYInit1 := float64(100), float64(100)
	Particule1.PositionX, Particule1.PositionY = PositionXInit1, PositionYInit1
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1

	Particule2 := Test.Basique_Particule()
	Particule2.ScaleX, Particule2.ScaleY = 1, 1
	SpeedXInit2, SpeedYInit2 := float64(10), float64(-2)
	PositionXInit2, PositionYInit2 := float64(200), float64(100)
	Particule2.PositionX, Particule2.PositionY = PositionXInit2, PositionYInit2
	Particule2.SpeedX, Particule2.SpeedY = SpeedXInit2, SpeedYInit2

	l.PushFront(&Particule2)
	l.PushFront(&Particule1)

	for NuméroParticule := l.Front(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Next() {
		ParticuleActuelle, ok := NuméroParticule.Value.(*particles.Particle)
		if ok {
			CollisionAmongParticle(NuméroParticule, ParticuleActuelle)
		}
	}

	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	particuleEnd2, _ := l.Front().Next().Value.(*particles.Particle)

	if particuleEnd1.SpeedX != SpeedXInit1 || particuleEnd1.SpeedY != SpeedYInit1 || particuleEnd2.SpeedX != SpeedXInit2 || particuleEnd2.SpeedY != SpeedYInit2 {
		t.Error("les vitesses des particules on été modifier alors qu'il n'y a pas eu de collision entre particules")
	} else if particuleEnd1.PositionX != PositionXInit1 || particuleEnd1.PositionY != PositionYInit1 || particuleEnd2.PositionX != PositionXInit2 || particuleEnd2.PositionY != PositionYInit2 {
		t.Error("les position des particules on été modifier alors qu'il n'y a pas eu de collision entre particules")
	}

}

func TestCollisionAmongParticleTrue1(t *testing.T) {

	l := list.New()

	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	PositionXInit1, PositionYInit1 := float64(100), float64(100)
	Particule1.PositionX, Particule1.PositionY = PositionXInit1, PositionYInit1
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1

	Particule2 := Test.Basique_Particule()
	Particule2.ScaleX, Particule2.ScaleY = 1, 1
	SpeedXInit2, SpeedYInit2 := float64(10), float64(-2)
	PositionXInit2, PositionYInit2 := float64(99.5), float64(100)
	Particule2.PositionX, Particule2.PositionY = PositionXInit2, PositionYInit2
	Particule2.SpeedX, Particule2.SpeedY = SpeedXInit2, SpeedYInit2

	l.PushFront(&Particule2)
	l.PushFront(&Particule1)

	for NuméroParticule := l.Front(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Next() {
		ParticuleActuelle, ok := NuméroParticule.Value.(*particles.Particle)
		if ok {
			CollisionAmongParticle(NuméroParticule, ParticuleActuelle)
		}
	}

	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	particuleEnd2, _ := l.Front().Next().Value.(*particles.Particle)

	if particuleEnd1.SpeedX != SpeedXInit2 || particuleEnd1.SpeedY != SpeedYInit2 || particuleEnd2.SpeedX != SpeedXInit1 || particuleEnd2.SpeedY != SpeedYInit1 {
		t.Error("les vitesses des particules ne c'est pas inverser avec la particule qu'elles ont toucher")
	} else if particuleEnd1.PositionX != PositionXInit1+0.1 || particuleEnd1.PositionY != PositionYInit1 || particuleEnd2.PositionX != PositionXInit2-0.1 || particuleEnd2.PositionY != PositionYInit2 {
		t.Error("les positions n'on pas été modifier correctement  alors que des particules ce sont chevauché")
	}

}

func TestCollisionAmongParticleTrue2(t *testing.T) {

	l := list.New()

	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	PositionXInit1, PositionYInit1 := float64(100), float64(100)
	Particule1.PositionX, Particule1.PositionY = PositionXInit1, PositionYInit1
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1

	Particule2 := Test.Basique_Particule()
	Particule2.ScaleX, Particule2.ScaleY = 1, 1
	SpeedXInit2, SpeedYInit2 := float64(10), float64(-2)
	PositionXInit2, PositionYInit2 := float64(100), float64(99.5)
	Particule2.PositionX, Particule2.PositionY = PositionXInit2, PositionYInit2
	Particule2.SpeedX, Particule2.SpeedY = SpeedXInit2, SpeedYInit2

	l.PushFront(&Particule2)
	l.PushFront(&Particule1)

	for NuméroParticule := l.Front(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Next() {
		ParticuleActuelle, ok := NuméroParticule.Value.(*particles.Particle)
		if ok {
			CollisionAmongParticle(NuméroParticule, ParticuleActuelle)
		}
	}

	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	particuleEnd2, _ := l.Front().Next().Value.(*particles.Particle)

	if particuleEnd1.SpeedX != SpeedXInit2 || particuleEnd1.SpeedY != SpeedYInit2 || particuleEnd2.SpeedX != SpeedXInit1 || particuleEnd2.SpeedY != SpeedYInit1 {
		t.Error("les vitesses des particules ne c'est pas inverser avec la particule qu'elles ont toucher")
	} else if particuleEnd1.PositionX != PositionXInit1 || particuleEnd1.PositionY != PositionYInit1+0.1 || particuleEnd2.PositionX != PositionXInit2 || particuleEnd2.PositionY != PositionYInit2-0.1 {
		t.Error("les positions n'on pas été modifier correctement  alors que des particules ce sont chevauché")
	}

}
