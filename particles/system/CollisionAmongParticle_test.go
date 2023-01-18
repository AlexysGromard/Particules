package system

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestCollisionAmongParticleFalse vérifie que les particules ne produisent pas de collision entre elles
func TestCollisionAmongParticleFalse(t *testing.T) {
	// Création de la liste de particules
	l := list.New()
	// Création de la première particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	PositionXInit1, PositionYInit1 := float64(100), float64(100)
	Particule1.PositionX, Particule1.PositionY = PositionXInit1, PositionYInit1
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Création de la deuxième particule
	Particule2 := Test.Basique_Particule()
	Particule2.ScaleX, Particule2.ScaleY = 1, 1
	SpeedXInit2, SpeedYInit2 := float64(10), float64(-2)
	PositionXInit2, PositionYInit2 := float64(200), float64(100)
	Particule2.PositionX, Particule2.PositionY = PositionXInit2, PositionYInit2
	Particule2.SpeedX, Particule2.SpeedY = SpeedXInit2, SpeedYInit2
	// Ajout des particules à la liste
	l.PushFront(&Particule2)
	l.PushFront(&Particule1)
	// Boucle sur la liste de particules
	// Pour chaque particule, on vérifie si elle est en collision avec une autre particule
	for NuméroParticule := l.Front(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Next() {
		ParticuleActuelle, ok := NuméroParticule.Value.(*particles.Particle)
		if ok {
			collisionAmongParticle(NuméroParticule, ParticuleActuelle)
		}
	}
	// On récupère les particules de la liste
	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	particuleEnd2, _ := l.Front().Next().Value.(*particles.Particle)
	// On vérifie que la vitesse des particules n'ont pas été modifier du fait d'une collision
	if particuleEnd1.SpeedX != SpeedXInit1 || particuleEnd1.SpeedY != SpeedYInit1 || particuleEnd2.SpeedX != SpeedXInit2 || particuleEnd2.SpeedY != SpeedYInit2 {
		t.Error("Les vitesses des particules ont été modifiées alors qu'il n'y a pas eu de collision entre particules")
		// On vérifie que la position des particules n'ont pas été modifié du fait d'une collision
	} else if particuleEnd1.PositionX != PositionXInit1 || particuleEnd1.PositionY != PositionYInit1 || particuleEnd2.PositionX != PositionXInit2 || particuleEnd2.PositionY != PositionYInit2 {
		t.Error("Les positions des particules ont été modifiées alors qu'il n'y a pas eu de collision entre particules")
	}

}

// Le test TestCollisionAmongParticleTrue1 vérifie que les particules produisent bien une collision entre elles
func TestCollisionAmongParticleTrue1(t *testing.T) {
	// Création de la liste de particules
	l := list.New()
	// Création de la première particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	PositionXInit1, PositionYInit1 := float64(100), float64(100)
	Particule1.PositionX, Particule1.PositionY = PositionXInit1, PositionYInit1
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Création de la deuxième particule
	Particule2 := Test.Basique_Particule()
	Particule2.ScaleX, Particule2.ScaleY = 1, 1
	SpeedXInit2, SpeedYInit2 := float64(10), float64(-2)
	PositionXInit2, PositionYInit2 := float64(99.5), float64(100)
	Particule2.PositionX, Particule2.PositionY = PositionXInit2, PositionYInit2
	Particule2.SpeedX, Particule2.SpeedY = SpeedXInit2, SpeedYInit2
	// Ajout des particules à la liste
	l.PushFront(&Particule2)
	l.PushFront(&Particule1)
	// Boucle sur la liste de particules
	for NuméroParticule := l.Front(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Next() {
		ParticuleActuelle, ok := NuméroParticule.Value.(*particles.Particle)
		if ok {
			collisionAmongParticle(NuméroParticule, ParticuleActuelle)
		}
	}
	// On récupère les particules de la liste
	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	particuleEnd2, _ := l.Front().Next().Value.(*particles.Particle)
	// On vérifie que les vitesses des particules se sont bien inversées
	if particuleEnd1.SpeedX != SpeedXInit2 || particuleEnd1.SpeedY != SpeedYInit2 || particuleEnd2.SpeedX != SpeedXInit1 || particuleEnd2.SpeedY != SpeedYInit1 {
		t.Error("Les vitesses ne se sont pas inversé avec la particule qu'elle a touchée")
		// On vérifie que les positions des particules se sont bien modifiées
	} else if particuleEnd1.PositionX != PositionXInit1+0.1 || particuleEnd1.PositionY != PositionYInit1 || particuleEnd2.PositionX != PositionXInit2-0.1 || particuleEnd2.PositionY != PositionYInit2 {
		t.Error("Les positions n'ont pas été modifiées correctement alors que des particules se sont chevauché")
	}

}

// Le test TestCollisionAmongParticleTrue2 vérifie que les particules produisent bien une collision entre elles
func TestCollisionAmongParticleTrue2(t *testing.T) {
	// Création de la liste de particules
	l := list.New()
	// Création de la première particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	PositionXInit1, PositionYInit1 := float64(100), float64(100)
	Particule1.PositionX, Particule1.PositionY = PositionXInit1, PositionYInit1
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Création de la deuxième particule
	Particule2 := Test.Basique_Particule()
	Particule2.ScaleX, Particule2.ScaleY = 1, 1
	SpeedXInit2, SpeedYInit2 := float64(10), float64(-2)
	PositionXInit2, PositionYInit2 := float64(100), float64(99.5)
	Particule2.PositionX, Particule2.PositionY = PositionXInit2, PositionYInit2
	Particule2.SpeedX, Particule2.SpeedY = SpeedXInit2, SpeedYInit2
	// Ajout des particules à la liste
	l.PushFront(&Particule2)
	l.PushFront(&Particule1)
	// Boucle sur la liste de particules
	for NuméroParticule := l.Front(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Next() {
		ParticuleActuelle, ok := NuméroParticule.Value.(*particles.Particle)
		if ok {
			collisionAmongParticle(NuméroParticule, ParticuleActuelle)
		}
	}
	// On récupère les particules de la liste
	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	particuleEnd2, _ := l.Front().Next().Value.(*particles.Particle)
	// On vérifie que les vitesses des particules se sont bien inversées
	if particuleEnd1.SpeedX != SpeedXInit2 || particuleEnd1.SpeedY != SpeedYInit2 || particuleEnd2.SpeedX != SpeedXInit1 || particuleEnd2.SpeedY != SpeedYInit1 {
		t.Error("Les vitesses ne se sont pas inversé avec la particule qu'elle a touchée")
		// On vérifie que les positions des particules se sont bien modifiées
	} else if particuleEnd1.PositionX != PositionXInit1 || particuleEnd1.PositionY != PositionYInit1+0.1 || particuleEnd2.PositionX != PositionXInit2 || particuleEnd2.PositionY != PositionYInit2-0.1 {
		t.Error("Les positions n'ont pas été modifiées correctement alors que des particules se sont chevauché")
	}

}
