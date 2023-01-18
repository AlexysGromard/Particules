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
	// On crée un systeme
	sys := System{Content: list.New(), DeadList: list.New()}
	// On crée une particule, avec une vie de 50 et une position en X et Y de 100
	Particule := Test.Basique_Particule()
	Particule.PositionX, Particule.PositionY = 100, 100
	Particule.Life = 50
	// On ajoute la particule au systeme
	sys.Content.PushFront(&Particule)
	// On met le parametre windowSizeX et Y à 200 et le parametre MarginOutsideScreen à 1
	config.General.WindowSizeX = 200
	config.General.WindowSizeY = 200
	config.General.MarginOutsideScreen = 1
	// On met à jour le systeme
	sys.Update()
	// On vérifie que la particule est bien supprimé
	if sys.Content.Len() != 1 {
		t.Error("La particule n'a pas été supprimé alors qu'elle est en dehors de l'écran")
	}
}

// Le test TestUpdateKillParticuleIfOutside2 vérifie que la particule est bien supprimé si elle est en dehors de l'écran
func TestUpdateKillParticuleIfOutside2(t *testing.T) {
	// On crée un systeme
	sys := System{Content: list.New(), DeadList: list.New()}
	// On crée une particule, avec une vie de 50 et une position en X et Y de 100
	Particule := Test.Basique_Particule()
	Particule.PositionX, Particule.PositionY = 300, 100
	Particule.Life = 50
	// On ajoute la particule au systeme
	sys.Content.PushFront(&Particule)
	// On met le parametre windowSizeX et Y à 200 et le parametre MarginOutsideScreen à 1
	config.General.WindowSizeX = 200
	config.General.WindowSizeY = 200
	config.General.MarginOutsideScreen = 1
	// On met à jour le systeme
	sys.Update()
	// On vérifie que la particule est bien supprimé
	if sys.Content.Len() != 0 {
		t.Error("La particule n'a pas été supprimé alors qu'elle est en dehors de l'écran")
	}
}

// Le test TestUpdateKillParticuleIfOutside3 vérifie que la particule est bien supprimé si elle est en dehors de l'écran
func TestUpdateKillParticuleIfOutside3(t *testing.T) {
	// On crée un systeme
	sys := System{Content: list.New(), DeadList: list.New()}
	// On crée une particule, avec une vie de 50 et une position en X et Y de 100
	Particule := Test.Basique_Particule()
	Particule.PositionX, Particule.PositionY = 100, 300
	Particule.Life = 50
	// On ajoute la particule au systeme
	sys.Content.PushFront(&Particule)
	// On met le parametre windowSizeX et Y à 200 et le parametre MarginOutsideScreen à 1
	config.General.WindowSizeX = 200
	config.General.WindowSizeY = 200
	config.General.MarginOutsideScreen = 1
	// On met à jour le systeme
	sys.Update()
	// On vérifie que la particule est bien supprimé
	if sys.Content.Len() != 0 {
		t.Error("La particule n'a pas été supprimé alors qu'elle est en dehors de l'écran")
	}
}

// Le test TestUpdateKillParticuleIf0Life vérifie que la particule est bien supprimé si elle a une vie de 0
func TestUpdateKillParticuleIf0Life(t *testing.T) {
	// On crée un systeme
	sys := System{Content: list.New(), DeadList: list.New()}
	// On crée une particule, avec une vie de 0 et une position en X et Y de 100
	Particule := Test.Basique_Particule()
	Particule.PositionX, Particule.PositionY = 100, 100
	Particule.Life = 0
	// On ajoute la particule au systeme
	sys.Content.PushFront(&Particule)
	// On met le parametre windowSizeX et Y à 200 et le parametre MarginOutsideScreen à 1
	config.General.WindowSizeX = 200
	config.General.WindowSizeY = 200
	config.General.MarginOutsideScreen = 1
	// On met à jour le systeme
	sys.Update()
	// On vérifie que la particule est bien supprimé
	if sys.Content.Len() != 0 {
		t.Error("La particule n'a pas été supprimé alors qu'elle a une vie de 0")
	}
}

func TestUpdateCollisionAmong(t *testing.T) {
	sys := System{Content: list.New(), DeadList: list.New()}

	Particule := Test.Basique_Particule()
	Particule.PositionX, Particule.PositionY = 100, 100
	Particule.Life = 0

	sys.Content.PushFront(&Particule)

	config.General.WindowSizeX = 200
	config.General.WindowSizeY = 200
	config.General.MarginOutsideScreen = 1

	sys.Update()

	if sys.Content.Len() != 0 {
		t.Error("La vie n'a pas correctement diminuer alors que le parametre HaveLife est à true")
	}
}
