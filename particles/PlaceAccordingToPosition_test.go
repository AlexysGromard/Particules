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

	Particle1.PositionX = 10
	Particle2.PositionX = 30
	Particle3.PositionX = 40

	Particle1.SpeedX = 20
	Particle2.SpeedX = -5
	Particle3.SpeedX = -3

	sys.Content.PushFront(&Particle3)
	sys.Content.PushFront(&Particle2)
	sys.Content.PushFront(&Particle1)

	for i := sys.Content.Front(); i != nil; i = i.Next() {
		p, ok := i.Value.(*Particle)
		if ok {
			p.PositionX = p.PositionX + p.SpeedX
		}
	}
	for i := sys.Content.Front(); i != nil; i = i.Next() {

		sys.PlaceAccordingToPosition(i)

	}

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
/*
func TestPlaceAccordingToPosition2(t *testing.T) {

	sys := System{Content: list.New(), DeadList: list.New()}

	Particle1 := Basique_Particule()
	Particle2 := Basique_Particule()
	Particle3 := Basique_Particule()
	Particle4 := Basique_Particule()
	Particle5 := Basique_Particule()


	Particle1.PositionX = 10
	Particle2.PositionX = 30
	Particle3.PositionX = 40

	Particle1.SpeedX = 20
	Particle2.SpeedX = -5
	Particle3.SpeedX = -3

	sys.Content.PushFront(&Particle3)
	sys.Content.PushFront(&Particle2)
	sys.Content.PushFront(&Particle1)

	for i := sys.Content.Front(); i != nil; i = i.Next() {
		p, ok := i.Value.(*Particle)
		if ok {
			p.PositionX = p.PositionX + p.SpeedX
		}
	}
	for i := sys.Content.Front(); i != nil; i = i.Next() {

		sys.PlaceAccordingToPosition(i)

	}

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
*/