package System

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

func TestPlaceAccordingToPosition1(t *testing.T) {

	sys := System{Content: list.New(), DeadList: list.New()}

	Particle1 := Test.Basique_Particule()
	Particle2 := Test.Basique_Particule()
	Particle3 := Test.Basique_Particule()

	Particle1.PositionX = 30
	Particle2.PositionX = 25
	Particle3.PositionX = 37

	sys.Content.PushFront(&Particle3)
	sys.Content.PushFront(&Particle2)
	sys.Content.PushFront(&Particle1)

	sys.PlaceAccordingToPosition()

	a := sys.Content.Front()
	b := a.Next()
	c := b.Next()

	pa, _ := a.Value.(*particles.Particle)
	pb, _ := b.Value.(*particles.Particle)
	pc, _ := c.Value.(*particles.Particle)

	if pa != &Particle2 && pb != &Particle1 && pc != &Particle3 {
		t.Error("le tri n'a pas bien fonctionner")
	}

}

func TestPlaceAccordingToPosition2(t *testing.T) {

	sys := System{Content: list.New(), DeadList: list.New()}

	Particle1 := Test.Basique_Particule()
	Particle2 := Test.Basique_Particule()
	Particle3 := Test.Basique_Particule()
	Particle4 := Test.Basique_Particule()
	Particle5 := Test.Basique_Particule()

	Particle1.PositionX = 27
	Particle2.PositionX = 31
	Particle3.PositionX = 4
	Particle4.PositionX = 42
	Particle5.PositionX = 29

	sys.Content.PushFront(&Particle5)
	sys.Content.PushFront(&Particle4)
	sys.Content.PushFront(&Particle3)
	sys.Content.PushFront(&Particle2)
	sys.Content.PushFront(&Particle1)

	sys.PlaceAccordingToPosition()

	a := sys.Content.Front()
	b := a.Next()
	c := b.Next()
	d := c.Next()
	e := d.Next()

	pa, _ := a.Value.(*particles.Particle)
	pb, _ := b.Value.(*particles.Particle)
	pc, _ := c.Value.(*particles.Particle)
	pd, _ := d.Value.(*particles.Particle)
	pe, _ := e.Value.(*particles.Particle)

	if pa != &Particle3 && pb != &Particle1 && pc != &Particle5 && pd != &Particle2 && pe != &Particle4 {
		t.Error("le tri n'a pas bien fonctionné", pa, pb, pc, pd, pe)
	}

}
