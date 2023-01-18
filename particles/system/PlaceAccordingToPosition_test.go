package system

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestPlaceAccordingToPosition vérifie que la fonction placeAccordingToPosition trie bien les particules
// en fonction de leur position
// Si ce n'est pas le cas, le test échoue
func TestPlaceAccordingToPosition(t *testing.T) {

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

	sys.placeAccordingToPosition()

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
