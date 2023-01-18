package installationParticles

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestInsertionAccordingToPositionXVide vérifie que la fonction InsertionAccordingToPositionX ajoute bien une particule à la liste
// Si la liste n'est pas égale à 1, le test échoue
func TestInsertionAccordingToPositionXVide(t *testing.T) {
	l := list.New()

	particule := Test.Basique_Particule()

	InsertionAccordingToPositionX(l, &particule, 100)

	if l.Len() != 1 {
		t.Error("la liste devrait contenir une particule")
	}
}

// Le test TestInsertionAccordingToPositionXUnElement1 vérifie que la fonction InsertionAccordingToPositionX ajoute bien 2 particule à la liste
// Si la liste n'est pas égale à 2, le test échoue
// Si les particules ne sont pas trié dans la bonne ordre, le test échoue également
func TestInsertionAccordingToPositionXUnElement1(t *testing.T) {
	l := list.New()

	particule1 := Test.Basique_Particule()
	particule1.PositionX = 10
	l.PushFront(&particule1)

	particule2 := Test.Basique_Particule()
	particule2.PositionX = 30
	InsertionAccordingToPositionX(l, &particule2, 100)

	if l.Len() != 2 {
		t.Error("la liste devrait contenir deux particules")
	}
	if l.Front().Value.(*particles.Particle) != &particule1 || l.Front().Next().Value.(*particles.Particle) != &particule2 {
		t.Error("les particules ne sont pas trie dans le bon ordre")
	}
}

// Le test TestInsertionAccordingToPositionXUnElement2 vérifie que la fonction InsertionAccordingToPositionX ajoute bien 2 particule à la liste
// Si la liste n'est pas égale à 2, le test échoue
// Si les particules ne sont pas trié dans la bonne ordre, le test échoue également
func TestInsertionAccordingToPositionXUnElement2(t *testing.T) {
	l := list.New()

	particule1 := Test.Basique_Particule()
	particule1.PositionX = 200
	l.PushFront(&particule1)

	particule2 := Test.Basique_Particule()
	particule2.PositionX = 30
	InsertionAccordingToPositionX(l, &particule2, 100)

	if l.Len() != 2 {
		t.Error("la liste devrait contenir deux particules")
	}
	if l.Front().Value.(*particles.Particle) != &particule2 || l.Front().Next().Value.(*particles.Particle) != &particule1 {
		t.Error("les particules ne sont pas trie dans le bon ordre")
	}
}

func TestInsertionAccordingToPositionXTroisElement(t *testing.T) {
	l := list.New()

	particule1 := Test.Basique_Particule()
	particule2 := Test.Basique_Particule()
	particule3 := Test.Basique_Particule()
	particule1.PositionX = 10
	particule2.PositionX = 100
	particule3.PositionX = 300
	l.PushFront(&particule3)
	l.PushFront(&particule2)
	l.PushFront(&particule1)

	particule4 := Test.Basique_Particule()
	particule4.PositionX = 150
	InsertionAccordingToPositionX(l, &particule4, 200)

	if l.Len() != 4 {
		t.Error("la liste devrait contenir deux particules")
	}
	if l.Front().Value.(*particles.Particle) != &particule1 || l.Front().Next().Value.(*particles.Particle) != &particule2 || l.Front().Next().Next().Value.(*particles.Particle) != &particule4 || l.Front().Next().Next().Next().Value.(*particles.Particle) != &particule3 {
		t.Error("les particules ne sont pas triées dans le bon ordre")
	}
}
