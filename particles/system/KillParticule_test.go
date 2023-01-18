package system

import (
	"container/list"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestKillParticule vérifie que la fonction KillParticule ajoute bien une particule à la liste des particules mortes
// et la supprime de la liste des particules du system
// Si deadList est vide et que la liste des particules est à 1, le test échoue
func TestKillParticule(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	sys.Content.PushFront(Test.Basique_Particule())

	sys.killParticule(sys.Content.Front())

	if sys.Content.Len() != 0 || sys.DeadList.Len() != 1 {
		t.Error("La liste contente devrait avoir 0 élément et Dead List 1 élément mais pour vous Content à ", sys.Content.Len(), " élément et Dead Liste à ", sys.DeadList.Len(), " élément")
	}
}
