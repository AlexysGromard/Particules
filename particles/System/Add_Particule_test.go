package System

import (
	"container/list"
	"testing"
	"project-particles/particles/Test"
)

func TestAdd_ParticuleFromNothing(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	sys.Add_Particule()

	if sys.Content.Len() != 1 {
		t.Error("il y a ", sys.Content.Len(), " particule dans le system ,mais il devrait en avoire que 1")
	}
	if sys.DeadList.Len() != 0{
		t.Error("il y a ", sys.DeadList.Len(), " particule dans la liste des particule morte, mais il devrait en avoire 0")
	}
}

func TestAdd_ParticuleFromDeadList(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	particule := Test.Basique_Particule()

	sys.Content.PushFront(&particule)
	sys.DeadList.PushFront(&particule)

	sys.Add_Particule()

	if sys.Content.Len() != 2{
		t.Error("il y a ", sys.Content.Len(), " particule dans le system ,mais il devrait en avoire que 1", )
	}
	if sys.DeadList.Len() != 0{
		t.Error("il y a ", sys.DeadList.Len(), " particule dans le system ,mais il devrait en avoire que 0")
	}
}
