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
	resetConfigUpdate()
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
	resetConfigUpdate()
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
	resetConfigUpdate()
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
	resetConfigUpdate()
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

// Le test TestUpdateKillParticuleIfOutsideFalse vérifie qu'il n'y pas de collision
func TestUpdateCollisionAmongParticleFalse1(t *testing.T) {
	if !verificationCollisionAmongParticle(false, true) {
		t.Error("il y a eu une collision alors qu'elles sont désactivé")
	}
}

// Le test TestUpdateCollisionAmongParticleFalse2 vérifie qu'il n'y pas de collision
func TestUpdateCollisionAmongParticleFalse2(t *testing.T) {
	if !verificationCollisionAmongParticle(true, false) {
		t.Error("il y a eu une collision alors qu'elles sont désactivé")
	}
}

// Le test TestUpdateCollisionAmongParticleTrue vérifie qu'il y a bien une collision
func TestUpdateCollisionAmongParticleTrue(t *testing.T) {
	if verificationCollisionAmongParticle(true, true) {
		t.Error("il n'y a pas eu de collision alors qu'elle sont activé")
	}
}

// Le test TestUpdateCollisionWithWallFalse1 vérifie qu'il n'y pas de collision avec un bord
func TestUpdateCollisionWithWallFalse1(t *testing.T) {
	if !verificationCollisionWithWall(false, true) {
		t.Error("il y a eu une collision avec un bord alors qu'elle sont désactivé")
	}
}

// Le test TestUpdateCollisionWithWallFalse2 vérifie qu'il n'y pas de collision avec un bord
func TestUpdateCollisionWithWallFalse2(t *testing.T) {
	if !verificationCollisionWithWall(true, false) {
		t.Error("il y a eu une collision avec un bord alors qu'elle sont désactivé")
	}
}

// Le test TestUpdateCollisionWithWallTrue vérifie qu'il y a bien une collision avec un bord
func TestUpdateCollisionWithWallTrue(t *testing.T) {
	if verificationCollisionWithWall(true, true) {
		t.Error("il n'y a pas eu de collision avec un bord alors qu'elle sont activé")
	}
}

// Le test TestUpdateAjouterParticule vérifie que le nombre de particule vivante est bien égale au SpawnRate
func TestUpdateAjouterParticule(t *testing.T) {
	resetConfigUpdate()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.SpawnRate = 20

	sys.Update()

	if sys.Content.Len() != 20 {
		t.Error("le SpawnRate est à 20 mais il n'y a ", sys.Content.Len(), " particule vivante")
	}
}

// Le test TestUpdateSpawnCenter vérifie que le point de Spawn est bien au centre de l'écran
func TestUpdateSpawnCenter(t *testing.T) {
	resetConfigUpdate()
	config.General.SpawnCenter = true
	config.General.WindowSizeX = 200
	config.General.WindowSizeY = 200
	config.General.SpawnX, config.General.SpawnY = 0, 0

	sys := System{Content: list.New(), DeadList: list.New()}

	sys.Update()

	if config.General.SpawnX != 100 || config.General.SpawnY != 100 {
		t.Error("le point de Spawn X et Y ne sont pas au centre de l'écan alors que SpawnCenter est True ")
	}

}

// Le test TestUpdateAccordingToPosition vérifie que les particules sont bien modifiés selon la position
// de la particule
func TestUpdateAccordingToPosition(t *testing.T) {
	resetConfigUpdate()
	sys := System{Content: list.New(), DeadList: list.New()}

	particule := Test.Basique_Particule()
	particule.PositionX, particule.PositionY = 0, 0
	particule.SpeedX, particule.SpeedY = 0, 0
	particule.ColorRed = 1
	particule.ScaleX = 1
	particule.Opacity = 1
	particule.Rotation = 1
	sys.Content.PushFront(&particule)

	config.General.MinColorRed, config.General.MaxColorRed = 0, 1
	config.General.ScaleX = 1
	config.General.Opacity = 1
	config.General.Rotation = 3.14

	config.General.ChangeColorAccordingTo = 1
	config.General.ChangeScaleAccordingTo = 1
	config.General.ChangeOpacityAccordingTo = 1
	config.General.ChangeRotationAccordingTo = 1

	sys.Update()

	if sys.Content.Front().Value.(*particles.Particle).ColorRed != 0 {
		t.Error("la couleur n'a pas été modifier alors que ChangeColorAccordingTo est à 1")
	} else if sys.Content.Front().Value.(*particles.Particle).ScaleX != 0 {
		t.Error("le ScaleX n'a pas été modifier alors que ChangeScaleAccordingTo est à 1")
	} else if sys.Content.Front().Value.(*particles.Particle).Opacity != 0 {
		t.Error("l'Opacity n'a pas été modifier alors que ChangeOpacityAccordingTo est à 1")
	} else if sys.Content.Front().Value.(*particles.Particle).Rotation != 0 {
		t.Error("la Rotation n'a pas été modifier alors que ChangeRotationAccordingTo est à 1")
	}

}

// Le test TestUpdatePlaceAccordingToPosition vérifie que les particules sont bien trié selon la position
func TestUpdatePlaceAccordingToPosition(t *testing.T) {
	resetConfigUpdate()

	sys := System{Content: list.New(), DeadList: list.New()}

	Particle1 := Test.Basique_Particule()
	Particle2 := Test.Basique_Particule()

	Particle1.PositionX = 30
	Particle2.PositionX = 25

	sys.Content.PushFront(&Particle2)
	sys.Content.PushFront(&Particle1)

	config.General.Collision = true
	config.General.CollisionAmongParticle = true

	sys.Update()

	a := sys.Content.Front()
	b := a.Next()

	pa, _ := a.Value.(*particles.Particle)
	pb, _ := b.Value.(*particles.Particle)

	if pa != &Particle2 && pb != &Particle1 {
		t.Error("il n'y a pas eu de tri alors que Collision et CollisionAmongParticle sont à True")
	}

}

// Le test TestUpdateAccordingToLife vérifie que les particules sont bien modifiés selon la vie de la particule
func TestUpdateAccordingToLife(t *testing.T) {
	resetConfigUpdate()
	sys := System{Content: list.New(), DeadList: list.New()}

	particule := Test.Basique_Particule()
	particule.Life = 25
	particule.LifeInit = 50
	particule.SpeedX, particule.SpeedY = 0, 0
	particule.ColorRed = 1
	particule.ScaleX = 1
	particule.Opacity = 1
	particule.Rotation = 1
	sys.Content.PushFront(&particule)

	config.General.MinColorRed, config.General.MaxColorRed = 0, 1
	config.General.ScaleX = 1
	config.General.Opacity = 1
	config.General.Rotation = 3.14

	config.General.ChangeColorAccordingTo = 2
	config.General.ChangeScaleAccordingTo = 2
	config.General.ChangeOpacityAccordingTo = 2
	config.General.ChangeRotationAccordingTo = 2

	sys.Update()

	if sys.Content.Front().Value.(*particles.Particle).ColorRed != 0.5 {
		t.Error("la couleur n'a pas été modifier alors que ChangeColorAccordingTo est à 2")
	} else if sys.Content.Front().Value.(*particles.Particle).ScaleX != 0.5 {
		t.Error("le ScaleX n'a pas été modifier alors que ChangeScaleAccordingTo est à 2")
	} else if sys.Content.Front().Value.(*particles.Particle).Opacity != 0.5 {
		t.Error("l'Opacity n'a pas été modifier alors que ChangeOpacityAccordingTo est à 2")
	} else if sys.Content.Front().Value.(*particles.Particle).Rotation != 1.57 {
		t.Error("la Rotation n'a pas été modifier alors que ChangeRotationAccordingTo est à 2")
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

	resetConfigUpdate()
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
	resetConfigUpdate()
	// Création de la liste de particules
	l := list.New()
	// Création de la première particule, scale 1, vitesse 30,2, position 100,100
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	PositionXInit1 := float64(99.5)
	Particule1.PositionX, Particule1.PositionY = PositionXInit1, 100
	Particule1.SpeedX, Particule1.SpeedY = 0, 0
	// Création de la deuxième particule, scale 1, vitesse 10,-2, position 99,100
	Particule2 := Test.Basique_Particule()
	Particule2.ScaleX, Particule2.ScaleY = 1, 1
	Particule2.PositionX, Particule2.PositionY = 100, 100
	Particule2.SpeedX, Particule2.SpeedY = 0, 0
	// Ajout des particules à la liste
	l.PushFront(&Particule2)
	l.PushFront(&Particule1)

	config.General.Collision = Collision
	config.General.CollisionAmongParticle = CollisionAmongParticle

	sys := System{Content: l, DeadList: list.New()}

	sys.Update()

	return sys.Content.Front().Value.(*particles.Particle).PositionX != PositionXInit1-0.1

}

func verificationCollisionWithWall(Collision, CollisionWithWall bool) bool {
	resetConfigUpdate()
	// Création de la liste de particules
	l := list.New()
	// Définire les limite de l'écan
	LimiteX, LimiteY := 200, 200 // Définire les limite de l'écan
	// Création de la particule
	Particule1 := Test.Basique_Particule()
	Particule1.ScaleX, Particule1.ScaleY = 1, 1
	SpeedXInit1, SpeedYInit1 := float64(30), float64(2)
	Particule1.PositionX, Particule1.PositionY = 170, 100
	Particule1.SpeedX, Particule1.SpeedY = SpeedXInit1, SpeedYInit1
	// Ajout de la particule à la liste
	l.PushFront(&Particule1)

	config.General.WindowSizeX, config.General.WindowSizeY = LimiteX, LimiteY
	config.General.Collision = Collision
	config.General.CollisionWithWall = CollisionWithWall

	sys := System{Content: l, DeadList: list.New()}

	sys.Update()

	// On vérifie que la vitesse de la particule a été modifier
	return sys.Content.Front().Value.(*particles.Particle).SpeedX != -SpeedXInit1

}

func resetConfigUpdate() {
	config.General.WindowSizeX, config.General.WindowSizeY = 200, 200

	config.General.SpawnCenter = false
	config.General.SpawnX, config.General.SpawnY = 0, 0
	config.General.SpawnRate = 0
	config.General.Rotation = 0
	config.General.ScaleX, config.General.ScaleY = 1, 1
	config.General.Opacity = 1
	config.General.ColorRed, config.General.ColorGreen, config.General.ColorBlue = 1, 1, 1
	config.General.Gravity = 0
	config.General.HaveLife = false
	config.General.RandomLife = false
	config.General.Life = 1

	config.General.MarginOutsideScreen = 0

	config.General.ChangeColorAccordingTo = 0
	config.General.MinColorRed = 0
	config.General.MinColorGreen = 0
	config.General.MinColorBlue = 0
	config.General.MaxColorRed = 0
	config.General.MaxColorGreen = 0
	config.General.MaxColorBlue = 0
	config.General.ChangeScaleAccordingTo = 0
	config.General.ChangeRotationAccordingTo = 0
	config.General.ChangeOpacityAccordingTo = 0

	config.General.Collision = false
	config.General.CollisionAmongParticle = false
	config.General.CollisionWithWall = false
	// Interaction
	config.General.Interaction = false
	config.General.FollowMouse = false
}
