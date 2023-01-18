package system

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestCollisionWithWallFalse vérifie que la particule ne touche pas les bords
func TestCollisionWithWallFalse(t *testing.T) {
	// Création de la liste de particules
	l := list.New()
	// Définire les limite de l'écan
	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan
	// Création de la particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 100, 100
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Ajout de la particule à la liste
	l.PushFront(&Particule1)
	// On vérifie que la particule ne touche pas les bords
	CollisionWithWall(&Particule1, LimiteX, LimiteY)
	// On récupère la particule
	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	// On vérifie que la vitesse de la particule n'a pas été modifier
	if particuleEnd1.SpeedX != SpeedXInit1 || particuleEnd1.SpeedY != SpeedYInit1 {
		t.Error("La vitesse de la particule a été modifiée alors qu'elle n'a pas touché les bords")
	}
}

// Le test TestCollisionWithWallTrueRIGHT vérifie que la particule touche le bord du haut
func TestCollisionWithWallTrueUP(t *testing.T) {
	// Création de la liste de particules
	l := list.New()
	// Définire les limite de l'écan
	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan
	// Création de la particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 100, 0
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Ajout de la particule à la liste
	l.PushFront(&Particule1)
	// On vérifie que la particule touche le bord du haut
	CollisionWithWall(&Particule1, LimiteX, LimiteY)
	// On récupère la particule
	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	// On vérifie que la vitesse de la particule a été modifier
	if particuleEnd1.SpeedX != SpeedXInit1 || particuleEnd1.SpeedY != -SpeedYInit1 || particuleEnd1.PositionX != 100 || particuleEnd1.PositionY != 0.1 {
		t.Error("la vitesse de la particule a été modifiée alors qu'elle n'a pas touché les bords ou la particule n'a pas été éloigné de la bordure")
	}
}

// Le test TestCollisionWithWallTrueLOW vérifie que la particule touche le bord du bas
func TestCollisionWithWallTrueDown(t *testing.T) {
	// Création de la liste de particules
	l := list.New()
	// Définire les limite de l'écan
	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan
	// Création de la particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 100, 200
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Ajout de la particule à la liste
	l.PushFront(&Particule1)
	// On vérifie que la particule touche le bord du bas
	CollisionWithWall(&Particule1, LimiteX, LimiteY)
	// On récupère la particule
	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	// On vérifie que la vitesse de la particule a été modifier
	if particuleEnd1.SpeedX != SpeedXInit1 || particuleEnd1.SpeedY != -SpeedYInit1 || particuleEnd1.PositionX != 100 || particuleEnd1.PositionY != 199.9 {
		t.Error("la vitesse de la particule a été modifiée alors qu'elle n'a pas touché les bords ou la particule n'a pas été éloigné de la bordure")
	}
}

// Le test TestCollisionWithWallTrueLEFT vérifie que la particule touche le bord de gauche
func TestCollisionWithWallTrueLEFT(t *testing.T) {
	// Création de la liste de particules
	l := list.New()
	// Définire les limite de l'écan
	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan
	// Création de la particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 0, 100
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Ajout de la particule à la liste
	l.PushFront(&Particule1)
	// On vérifie que la particule touche le bord de gauche
	CollisionWithWall(&Particule1, LimiteX, LimiteY)
	// On récupère la particule
	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	// On vérifie que la vitesse de la particule a été modifier
	if particuleEnd1.SpeedX != -SpeedXInit1 || particuleEnd1.SpeedY != SpeedYInit1 || particuleEnd1.PositionX != 0.1 || particuleEnd1.PositionY != 100 {
		t.Error("La vitesse de la particule a été modifiée alors qu'elle n'a pas touché les bords ou la particule n'a pas été éloigné de la bordure")
	}
}

// Le test TestCollisionWithWallTrueRIGHT vérifie que la particule touche le bord de droite
func TestCollisionWithWallTrueRIGHT(t *testing.T) {
	// Création de la liste de particules
	l := list.New()
	// Définire les limite de l'écan
	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan
	// Création de la particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 200, 100
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Ajout de la particule à la liste
	l.PushFront(&Particule1)
	// On vérifie que la particule touche le bord de droite
	CollisionWithWall(&Particule1, LimiteX, LimiteY)
	// On récupère la particule
	particuleEnd1, _ := l.Front().Value.(*particles.Particle)
	// On vérifie que la vitesse de la particule a été modifier
	if particuleEnd1.SpeedX != -SpeedXInit1 || particuleEnd1.SpeedY != SpeedYInit1 || particuleEnd1.PositionX != 199.9 || particuleEnd1.PositionY != 100 {
		t.Error("La vitesse de la particule a été modifiée alors qu'elle n'a pas touché les bords ou la particule n'a pas été éloigné de la bordure")
	}
}
