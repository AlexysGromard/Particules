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

	elementa := sys.Content.Front()
	elementb := elementa.Next()
	elementc := elementb.Next()
	elementd := elementc.Next()
	elemente := elementd.Next()

	particulea, _ := elementa.Value.(*particles.Particle)
	particuleb, _ := elementb.Value.(*particles.Particle)
	particulec, _ := elementc.Value.(*particles.Particle)
	particuled, _ := elementd.Value.(*particles.Particle)
	particulee, _ := elemente.Value.(*particles.Particle)

	if particulea != &Particle3 && particuleb != &Particle1 && particulec != &Particle5 && particuled != &Particle2 && particulee != &Particle4 {
		t.Error("le tri n'a pas bien fonctionné")
	}

}
