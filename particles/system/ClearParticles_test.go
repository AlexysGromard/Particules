package system

import (
	"container/list"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestClearParticles0 vérifie que toutes les particules vivantes du system sont tuées
func TestClearParticles0(t *testing.T) {
	// Création du system
	Sys := System{Content: list.New(), DeadList: list.New()}
	// Ajout de 100 particules vivantes au system
	for NuméroToure := 0; NuméroToure < 100; NuméroToure++ {
		Sys.Content.PushFront(Test.Basique_Particule())
	}
	// On tue toutes les particules vivantes du system
	Sys.ClearParticles()
	// On vérifie que toutes les particules vivantes du system ont été tuées
	if Sys.Content.Len() != 0 {
		t.Error("Toutes les particules vivantes du system n'ont pas été tuées, il reste ", Sys.Content.Len(), " particule(s) vivante(s)")
	}
}
