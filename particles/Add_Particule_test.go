package particles

import (
	"container/list"
	"testing"
)

func TestAdd_ParticuleFromNothing(t *testing.T) {
	sys := System{Content: list.New()}
	DeadList := list.New()

	sys.Add_Particule(DeadList)

	if sys.Content.Len() != 1 {
		t.Error("il y a ", sys.Content.Len(), " particule dans le system ,mais il devrait en avoire que 1")
	}
	if DeadList.Len() != 0{
		t.Error("il y a ", DeadList.Len(), " particule dans la liste des particule morte, mais il devrait en avoire 0")
	}
}

func TestAdd_ParticuleFromDeadList(t *testing.T) {
	sys := System{Content: list.New()}
	DeadList := list.New()

	particule := Basique_Particule()

	sys.Content.PushFront(&particule)
	DeadList.PushFront(&particule)

	sys.Add_Particule(DeadList)

	if sys.Content.Len() != 1{
		t.Error("il y a ", sys.Content.Len(), " particule dans le system ,mais il devrait en avoire que 1")
	}
	if DeadList.Len() != 0{
		t.Error("il y a ", DeadList.Len(), " particule dans le system ,mais il devrait en avoire que 0")
	}
}
