package system

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestUpdateGravity vérifie que la gravité est bien appliqué
// Elle vérifie aussi que la vitesse en X n'est pas modifier
func TestUpdateGravity(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	Particule := Test.Basique_Particule()
	SpeedXInit, SpeedYInit := float64(20), float64(20)
	Particule.PositionX, Particule.PositionY = 100, 100
	Particule.SpeedX, Particule.SpeedY = SpeedXInit, SpeedYInit

	sys.Content.PushFront(&Particule)

	config.General.Gravity = 1
	config.General.WindowSizeX, config.General.WindowSizeY = 1000, 1000

	sys.Update()

	particuleEnd, _ := sys.Content.Front().Value.(*particles.Particle)

	if particuleEnd.SpeedX != SpeedXInit {
		t.Error("la vitesse en X à été modifier mais elle doit reter la même")
	} else if particuleEnd.SpeedY == SpeedYInit {
		t.Error("")
	} else if particuleEnd.SpeedY != SpeedYInit+config.General.Gravity {
		t.Error("la Gravite n'a pas été appliqué correctement")
	}
}

// Le test TestUpdateVitesse vérifie que la vitesse est bien appliqué
// Elle vérifie aussi que la position en X et Y a bien été modifier
// et que la vitesse a bien été modifié.
func TestUpdateVitesse(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	Particule := Test.Basique_Particule()
	Particule.SpeedX, Particule.SpeedY = 20, 10
	PositionXInit, PositionYInit := float64(80), float64(80)
	Particule.PositionX, Particule.PositionY = PositionXInit, PositionYInit

	sys.Content.PushFront(&Particule)

	config.General.Gravity = 0
	config.General.WindowSizeX, config.General.WindowSizeY = 1000, 1000

	sys.Update()

	particuleEnd := sys.Content.Front().Value.(*particles.Particle)

	if particuleEnd.PositionX == PositionXInit || particuleEnd.PositionY == PositionYInit {
		t.Error("la vitesse n'a pas été appliqué")
	} else if particuleEnd.PositionX != 100 || particuleEnd.PositionY != 90 {
		t.Error("la vitesse n'a pas été correctement appliqué")
	}

}

// Le Test TestUpdateLifeFalse vérifie que la vie n'est pas modifié si le parametre HaveLife est à false
func TestUpdateLifeFalse(t *testing.T) {
	// On crée un systeme
	sys := System{Content: list.New(), DeadList: list.New()}
	// On crée une particule, avec une vie de 50
	Particule := Test.Basique_Particule()
	LifeStart := 50
	Particule.Life = LifeStart
	// On ajoute la particule au systeme
	sys.Content.PushFront(&Particule)
	// On met le parametre HaveLife à false
	config.General.HaveLife = false
	// On met à jour le systeme
	sys.Update()
	// On récupère la particule
	particuleEnd := sys.Content.Front().Value.(*particles.Particle)
	// On vérifie que la vie n'a pas été modifié
	if particuleEnd.Life != LifeStart {
		t.Error("La vie à été modifier alors que le parametre haveLife est à false")
	}
}

// Le test TestUpdateLifeTrue vérifie que la vie est bien modifié si le parametre HaveLife est à true
func TestUpdateLifeTrue(t *testing.T) {
	// On crée un systeme
	sys := System{Content: list.New(), DeadList: list.New()}
	// On crée une particule, avec une vie de 50
	Particule := Test.Basique_Particule()
	LifeStart := 50
	Particule.Life = LifeStart
	// On ajoute la particule au systeme
	sys.Content.PushFront(&Particule)
	// On met le parametre HaveLife à true
	config.General.HaveLife = true
	// On met à jour le systeme
	sys.Update()
	// On récupère la particule
	particuleEnd := sys.Content.Front().Value.(*particles.Particle)
	// On vérifie que la vie a bien été modifié
	if particuleEnd.Life != LifeStart-1 {
		t.Error("La vie n'a pas correctement diminuer alors que le parametre HaveLife est à true")
	}
}

// Le test TestUpdateKillParticuleIfOutside1 vérifie que la particule est bien supprimé si elle est en dehors de l'écran
// avec un parametre MarginOutsideScreen à 1
func TestUpdateKillParticuleIfOutside1(t *testing.T) {

	if verificationKillParticule(100, 100, 50, 1) {
		t.Error("La particule n'a pas été supprimé alors qu'elle est en dehors de l'écran")
	}
}

// Le test TestUpdateKillParticuleIfOutside2 vérifie que la particule est bien supprimé si elle est en dehors de l'écran
func TestUpdateKillParticuleIfOutside2(t *testing.T) {
	if verificationKillParticule(300, 100, 50, 0) {
		t.Error("La particule n'a pas été supprimé alors qu'elle est en dehors de l'écran")
	}
}

// Le test TestUpdateKillParticuleIfOutside3 vérifie que la particule est bien supprimé si elle est en dehors de l'écran
func TestUpdateKillParticuleIfOutside3(t *testing.T) {
	if verificationKillParticule(100, 300, 50, 0) {
		t.Error("La particule n'a pas été supprimé alors qu'elle est en dehors de l'écran")
	}
}

// Le test TestUpdateKillParticuleIf0Life vérifie que la particule est bien supprimé si elle a une vie de 0
func TestUpdateKillParticuleIf0Life(t *testing.T) {
	if verificationKillParticule(100, 100, 0, 0) {
		t.Error("La particule n'a pas été supprimé alors qu'elle a une vie de 0")
	}
}

func TestUpdateCollisionAmongParticleFalse1(t *testing.T) {
	if !verificationCollisionAmongParticle(false, true) {
		t.Error("il y a eu une collision alors qu'elle sont désactivé")
	}
}
func TestUpdateCollisionAmongParticleFalse2(t *testing.T) {
	if !verificationCollisionAmongParticle(true, false) {
		t.Error("il y a eu une collision alors qu'elle sont désactivé")
	}
}
func TestUpdateCollisionAmongParticleTrue(t *testing.T) {
	if verificationCollisionAmongParticle(true, true) {
		t.Error("il n'y a pas eu d collision alors qu'elle sont activé")
	}
}

// /
func TestUpdateCollisionWithWallFalse1(t *testing.T) {
	if !verificationCollisionWithWall(false, true) {
		t.Error("il y a eu une collision alors qu'elle sont désactivé")
	}
}
func TestUpdateCollisionWithWallFalse2(t *testing.T) {
	if !verificationCollisionWithWall(true, false) {
		t.Error("il y a eu une collision alors qu'elle sont désactivé")
	}
}
func TestUpdateCollisionWithWallTrue(t *testing.T) {
	if verificationCollisionWithWall(true, true) {
		t.Error("il n'y a pas eu d collision alors qu'elle sont activé")
	}
}

//
//
//
//
//
//
//
//
//

func verificationKillParticule(positionX, positionY float64, life, ParticuleFin int) bool {
	// On crée un systeme
	sys := System{Content: list.New(), DeadList: list.New()}
	// On crée une particule, avec une vie de 50 et une position en X et Y de 100
	Particule := Test.Basique_Particule()
	Particule.PositionX, Particule.PositionY = positionX, positionY
	Particule.Life = life
	// On ajoute la particule au systeme
	sys.Content.PushFront(&Particule)
	// On met le parametre windowSizeX et Y à 200 et le parametre MarginOutsideScreen à 1
	config.General.WindowSizeX = 200
	config.General.WindowSizeY = 200
	config.General.MarginOutsideScreen = 1
	// On met à jour le systeme
	sys.Update()
	// On vérifie que la particule est bien supprimé

	return sys.Content.Len() != ParticuleFin //t.Error("La particule n'a pas été supprimé alors qu'elle est en dehors de l'écran")

}

func verificationCollisionAmongParticle(Collision, CollisionAmongParticle bool) bool {
	// Création de la liste de particules
	l := list.New()
	// Création de la première particule, scale 1, vitesse 30,2, position 100,100
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	PositionXInit1 := float64(100)
	Particule1.PositionX, Particule1.PositionY = PositionXInit1, 2
	Particule1.SpeedX, Particule1.SpeedY = 100, 100
	// Création de la deuxième particule, scale 1, vitesse 10,-2, position 99,100
	Particule2 := Test.Basique_Particule()
	Particule2.ScaleX, Particule2.ScaleY = 1, 1
	Particule2.PositionX, Particule2.PositionY = 10, -2
	Particule2.SpeedX, Particule2.SpeedY = 99.5, 100
	// Ajout des particules à la liste
	l.PushFront(&Particule2)
	l.PushFront(&Particule1)

	config.General.Collision = Collision
	config.General.CollisionAmongParticle = CollisionAmongParticle

	sys := System{Content: l, DeadList: list.New()}

	sys.Update()

	return sys.Content.Front().Value.(*particles.Particle).PositionX != PositionXInit1+0.1

}

func verificationCollisionWithWall(Collision, CollisionWithWall bool) bool {

}
