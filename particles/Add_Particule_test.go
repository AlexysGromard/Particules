package particles

import (
	"container/list"
	"testing"
)

func TestAdd_ParticuleFromNothing(t *testing.T) {
	sys := System{Content: list.New()}

	sys.Add_Particule()

	if sys.Content.Len() != 1 {
		t.Error("il y a ", sys.Content.Len(), " particule dans le system ,mais il devrait en avoir que 1")
	}

}
