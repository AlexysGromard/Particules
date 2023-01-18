package System

import (
	"container/list"
	"project-particles/particles/Test"
	"testing"
)

func TestClearParticles0(t *testing.T) {
	Sys := System{Content: list.New(), DeadList: list.New()}

	for NuméroToure := 0; NuméroToure < 100; NuméroToure++ {
		Sys.Content.PushFront(Test.Basique_Particule())
	}

	Sys.ClearParticles()

	if Sys.Content.Len() != 0 {
		t.Error("Tout les Particule Vivante du System n'on pas été tuer il reste ", Sys.Content.Len(), " Particule")
	}
}
