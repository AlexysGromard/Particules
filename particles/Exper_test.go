package particles

import (
	"container/list"
	"fmt"
	"testing"
)

func TestDeplace(t *testing.T) {
	l := list.New()

	Particle1 := Basique_Particule()
	Particle2 := Basique_Particule()
	Particle3 := Basique_Particule()

	l.PushFront(&Particle1)
	l.PushFront(&Particle2)
	l.PushFront(&Particle3)
	fmt.Println("test ", l)

	p, ok := l.Front().Value.(*Particle)
	if ok {
		fmt.Println("test ", l)
		p.PositionX = 100
	}

	fmt.Println("test ", l)

	l.MoveAfter(l.Front(), l.Front().Next())

	fmt.Println(l)

	p, ok = l.Front().Value.(*Particle)
	if ok {
		fmt.Println("test ", l)
		if p.PositionX == 100 {
			t.Error("pas bien")
		}
	}

	p, ok = l.Front().Next().Value.(*Particle)
	if ok {
		fmt.Println("test ", l)
		if p.PositionX != 100 {
			t.Error("bien")
		}
	}

}
