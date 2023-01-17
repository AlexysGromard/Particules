package System

import (
	"container/list"
	"testing"
	"project-particles/particles/Test"

)

func TestResurrectParticule(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	sys.DeadList.PushFront(Test.Basique_Particule())

	sys.ResurrectParticule(sys.DeadList.Front(), 100, true, true)

	if sys.Content.Len() != 1 || sys.DeadList.Len() != 0 {
		t.Error("la liste Content devrait avoir 1 élément et DeadList 0 élément mais pour vous Content à ", sys.Content.Len(), " élément et DeadListe à ", sys.DeadList.Len(), " élément")
	}
}

func TestResurrectParticule2(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	sys.DeadList.PushFront(Test.Basique_Particule())

	sys.ResurrectParticule(sys.DeadList.Front(), 100, true, false)

	if sys.Content.Len() != 1 || sys.DeadList.Len() != 0 {
		t.Error("la liste Content devrait avoir 1 élément et DeadList 0 élément mais pour vous Content à ", sys.Content.Len(), " élément et DeadListe à ", sys.DeadList.Len(), " élément")
	}
}
