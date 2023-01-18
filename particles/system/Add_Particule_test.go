package system

import (
	"container/list"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestAdd_ParticuleFromNothing vérifie que la fonction Add_Particule ajoute bien une particule au system si la liste des particules mortes est vide
func TestAdd_ParticuleFromNothing(t *testing.T) {
	// Création du system
	sys := System{Content: list.New(), DeadList: list.New()}
	// Ajout d'une particule au system
	sys.Add_Particule()
	// On vérifie que la particule a bien été ajoutée au system
	if sys.Content.Len() != 1 {
		t.Error("il y a ", sys.Content.Len(), " particules dans le systeme ,mais il devrait y en avoir que 1")
	}
	// On vérifie que la liste des particules mortes est bien vide
	if sys.DeadList.Len() != 0 {
		t.Error("il y a ", sys.DeadList.Len(), " particule(s) dans la liste des particule morte, mais il devrait y en avoir 0")
	}
}

// Le test TestAdd_ParticuleFromDeadList vérifie que la fonction Add_Particule ajoute bien une particule au system à partir de la liste des particules mortes
// et que la liste des particules mortes est bien vide après l'ajout
// (Lorsque la liste des particules mortes n'est pas vide, la fonction Add_Particule prend une particule de la liste des particules mortes)
func TestAdd_ParticuleFromDeadList(t *testing.T) {
	// Création du system
	sys := System{Content: list.New(), DeadList: list.New()}
	// Ajout d'une particule au system
	particule := Test.Basique_Particule()
	// On tue la particule
	sys.Content.PushFront(&particule)
	sys.DeadList.PushFront(&particule)
	// On ajoute une nouvelle particule au systeme
	sys.Add_Particule()
	// On vérifie qu'il y ai bien une particule dans le systeme
	if sys.Content.Len() != 2 {
		t.Error("il y a ", sys.Content.Len(), " particules dans le systeme ,mais il devrait y avoir que 1")
	}
	// On vérifie que la liste des particules mortes est bien vide
	if sys.DeadList.Len() != 0 {
		t.Error("il y a ", sys.DeadList.Len(), " particule(s) dans le systeme ,mais il devrait y en avoir que 0")
	}
}
