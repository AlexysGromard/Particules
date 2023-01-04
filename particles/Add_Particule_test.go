package particles

import (
	"container/list"
	"testing"
)

func TestAdd_Particule(t *testing.T) {
	sys := System{Content: list.New()}
	DeadList := list.New()

	sys.Add_Particule(DeadList)

	if sys.Content.Len() != 1 {
		t.Error("il y a ", sys.Content.Len(), " particule dans le system ,mais il devrait en avoire que 1")
	}
}
