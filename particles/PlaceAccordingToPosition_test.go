package particles

import (
	"container/list"
	"testing"
)

func TestPlaceAccordingToPosition1(t *testing.T) {

	sys := System{Content: list.New(), DeadList: list.New()}

	Particle1 := Basique_Particule()
	Particle2 := Basique_Particule()
	Particle3 := Basique_Particule()

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

	pa, _ := a.Value.(*Particle)
	pb, _ := b.Value.(*Particle)
	pc, _ := c.Value.(*Particle)

	if pa != &Particle2 && pb != &Particle1 && pc != &Particle3 {
		t.Error("le tri n'a pas bien fonctionner")
	}

}

func TestPlaceAccordingToPosition2(t *testing.T) {

	sys := System{Content: list.New(), DeadList: list.New()}

	Particle1 := Basique_Particule()
	Particle2 := Basique_Particule()
	Particle3 := Basique_Particule()
	Particle4 := Basique_Particule()
	Particle5 := Basique_Particule()

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
	d := b.Next()
	e := b.Next()

	pa, _ := a.Value.(*Particle)
	pb, _ := b.Value.(*Particle)
	pc, _ := c.Value.(*Particle)
	pd, _ := d.Value.(*Particle)
	pe, _ := e.Value.(*Particle)

	if pa != &Particle3 && pb != &Particle1 && pc != &Particle5 && pd != &Particle2 && pe != &Particle4 {
		t.Error("le tri n'a pas bien fonctionner")
	}

}
