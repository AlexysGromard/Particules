package particles

import (
	"container/list"
	"fmt"
	"testing"
)

func TestPlaceAccordingToPosition(t *testing.T) {

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

	Particle1.ColorRed = 0.5
	Particle2.ColorGreen = 0.5
	Particle3.ColorBlue = 0.5

	sys.Content.PushFront(&Particle3)
	sys.Content.PushFront(&Particle2)
	sys.Content.PushFront(&Particle1)

	fmt.Println("taille de la liste ", sys.Content)

	fmt.Println(" 1 : ", Particle1.ColorRed, "/", Particle1.ColorGreen, "/", Particle1.ColorBlue, "; 2 : ", Particle2.ColorRed, "/", Particle2.ColorGreen, "/", Particle2.ColorBlue, "; 3 : ", Particle3.ColorRed, "/", Particle3.ColorGreen, "/", Particle3.ColorBlue)

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

	fmt.Println(" 1 : ", pa.ColorRed, "/", pa.ColorGreen, "/", pa.ColorBlue, "; 2 : ", pb.ColorRed, "/", pb.ColorGreen, "/", pb.ColorBlue, "; 3 : ", pc.ColorRed, "/", pc.ColorGreen, "/", pc.ColorBlue)
	if pa.ColorGreen != 0.5 && pb.ColorRed != 0.5 && pc.ColorBlue != 0.5 {
		t.Error("erruer")
	}

}
