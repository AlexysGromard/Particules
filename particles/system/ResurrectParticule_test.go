package system

import (
	"container/list"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestResurrectParticule vérifie que la fonction resurrectParticule ajoute bien une particule à la liste des particules du system
// et la supprime de la liste des particules mortes
// Si ça n'est pas le cas, le test échoue
func TestResurrectParticule(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	sys.DeadList.PushFront(Test.Basique_Particule())

	sys.resurrectParticule(sys.DeadList.Front(), 100, true, true)

	if sys.Content.Len() != 1 || sys.DeadList.Len() != 0 {
		t.Error("la liste Content devrait avoir 1 élément et DeadList 0 élément mais pour vous Content à ", sys.Content.Len(), " élément et DeadListe à ", sys.DeadList.Len(), " élément")
	}
}

// Le test TestResurrectParticule2 vérifie que la fonction resurrectParticule ajoute bien une particule à la liste des particules du system
// et la supprime de la liste des particules mortes
// Si ça n'est pas le cas, le test échoue
func TestResurrectParticule2(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	sys.DeadList.PushFront(Test.Basique_Particule())

	sys.resurrectParticule(sys.DeadList.Front(), 100, true, false)

	if sys.Content.Len() != 1 || sys.DeadList.Len() != 0 {
		t.Error("la liste Content devrait avoir 1 élément et DeadList 0 élément mais pour vous Content à ", sys.Content.Len(), " élément et DeadListe à ", sys.DeadList.Len(), " élément")
	}
}
