package particles

import (
	"testing"
	"container/list"
)

func TestKillParticule(t *testing.T){
	sys := System{Content: list.New(), DeadList: list.New()}

	sys.Content.PushFront(Basique_Particule())

	sys.KillParticule(sys.Content.Front())

	if sys.Content.Len() !=0 || sys.DeadList.Len() !=1{
		t.Error("la liste content devrait avoir 0 élément et DeadList 1 élément mais pour vous Content à ",sys.Content.Len()," élément et DeadListe à ",sys.DeadList.Len()," élément")
	}
}