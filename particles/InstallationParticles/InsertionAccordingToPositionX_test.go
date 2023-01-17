package InstallationParticles

import (
	"testing"
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
)



func TestInsertionAccordingToPositionXVide(t *testing.T){
	l := list.New()

	particule := Test.Basique_Particule()

	InsertionAccordingToPositionX(l,&particule,100)

	if l.Len()!= 1{
		t.Error("la liste devrais contenire une particule")
	}
}

func TestInsertionAccordingToPositionXUnElement1(t *testing.T){
	l := list.New()

	particule1 := Test.Basique_Particule() 
	particule1.PositionX = 10
	l.PushFront(&particule1)

	particule2 := Test.Basique_Particule() 
	particule2.PositionX = 30
	InsertionAccordingToPositionX(l,&particule2,100)


	if l.Len()!= 2{
		t.Error("la liste devrais contenire deux particule")
	}
	if l.Front().Value.(*particles.Particle) != &particule1 || l.Front().Next().Value.(*particles.Particle) != &particule2{
		t.Error("les paticule ne sont pas trie dans la bonne ordre")
	}
}

func TestInsertionAccordingToPositionXUnElement2(t *testing.T){
	l := list.New()

	particule1 := Test.Basique_Particule() 
	particule1.PositionX = 10
	l.PushFront(&particule1)

	particule2 := Test.Basique_Particule() 
	particule2.PositionX = 200
	InsertionAccordingToPositionX(l,&particule2,100)


	if l.Len()!= 2{
		t.Error("la liste devrais contenire deux particule")
	}
	if l.Front().Value.(*particles.Particle) != &particule1 || l.Front().Next().Value.(*particles.Particle) != &particule2{
		t.Error("les paticule ne sont pas trie dans la bonne ordre")
	}
}

func TestInsertionAccordingToPositionX2DeuxElement(t *testing.T){
	l := list.New()

	particule1 := Test.Basique_Particule() 
	particule2 := Test.Basique_Particule() 
	particule1.PositionX = 10
	particule2.PositionX = 300
	l.PushFront(&particule2)
	l.PushFront(&particule1)

	particule3 := Test.Basique_Particule() 
	particule3.PositionX = 150
	InsertionAccordingToPositionX(l,&particule3,200)

	if l.Len()!= 3{
		t.Error("la liste devrais contenire deux particule")
	}
	if l.Front().Value.(*particles.Particle) != &particule1 || l.Front().Next().Value.(*particles.Particle) != &particule3 || l.Front().Next().Next().Value.(*particles.Particle) != &particule2{
		t.Error("les paticules ne sont pas trié dans la bonne ordre")
	}
}

func TestInsertionAccordingToPositionX2TroisElement(t *testing.T){
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
	InsertionAccordingToPositionX(l,&particule4,200)

	if l.Len()!= 4{
		t.Error("la liste devrais contenire deux particule")
	}
	if l.Front().Value.(*particles.Particle) != &particule1 || l.Front().Next().Value.(*particles.Particle) != &particule2 || l.Front().Next().Next().Value.(*particles.Particle) != &particule4 || l.Front().Next().Next().Next().Value.(*particles.Particle)  != &particule3 {
		t.Error("les paticules ne sont pas trié dans la bonne ordre")
	}
}