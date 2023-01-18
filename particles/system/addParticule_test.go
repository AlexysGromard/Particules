package system

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"project-particles/particles/Test"
	"testing"
)

// Le test TestAdd_ParticuleFromNothing vérifie que la fonction Add_Particule ajoute bien une particule au system si la liste des particules mortes est vide
func TestAddParticuleFromNothing(t *testing.T) {
	// Création du system
	sys := System{Content: list.New(), DeadList: list.New()}
	// Ajout d'une particule au system
	sys.AddParticule()
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
func TestAddParticuleFromDeadList(t *testing.T) {
	// Création du system
	sys := System{Content: list.New(), DeadList: list.New()}
	// Ajout d'une particule au system
	particule := Test.Basique_Particule()
	// On tue la particule
	sys.Content.PushFront(&particule)
	sys.DeadList.PushFront(&particule)
	// On ajoute une nouvelle particule au systeme
	sys.AddParticule()
	// On vérifie qu'il y ai bien une particule dans le systeme
	if sys.Content.Len() != 2 {
		t.Error("il y a ", sys.Content.Len(), " particules dans le systeme ,mais il devrait y avoir que 1")
	}
	// On vérifie que la liste des particules mortes est bien vide
	if sys.DeadList.Len() != 0 {
		t.Error("il y a ", sys.DeadList.Len(), " particule(s) dans le systeme ,mais il devrait y en avoir que 0")
	}
}

func TestAddParticuleRandomSpawn(t *testing.T) {
	resetConfigAddParticule()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.RandomSpawn = true

	sys.AddParticule()
	sys.AddParticule()

	if sys.Content.Front().Value.(*particles.Particle).PositionX == sys.Content.Front().Next().Value.(*particles.Particle).PositionX && sys.Content.Front().Value.(*particles.Particle).PositionY == sys.Content.Front().Next().Value.(*particles.Particle).PositionY {
		t.Error("la possiton n'est pas aléatoire alors que RandomSpawn est à true ")
	}
}

func TestAddParticuleSpawnOnAnObject(t *testing.T) {
	resetConfigAddParticule()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.SpawnOnAnObject = true
	config.General.SpawnObject = "circle"
	config.General.SpawnObjectWidth = 500

	sys.AddParticule()
	sys.AddParticule()

	if sys.Content.Front().Value.(*particles.Particle).PositionX == sys.Content.Front().Next().Value.(*particles.Particle).PositionX && sys.Content.Front().Value.(*particles.Particle).PositionY == sys.Content.Front().Next().Value.(*particles.Particle).PositionY {
		t.Error("la possiton n'est pas aléatoire alors que SpawnOnAnObject est à true ")
	}
}

func TestAddParticuleNormalSpawn(t *testing.T) {
	resetConfigAddParticule()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.SpawnX = 100
	config.General.SpawnY = 110

	sys.AddParticule()

	if sys.Content.Front().Value.(*particles.Particle).PositionX != 100 || sys.Content.Front().Value.(*particles.Particle).PositionY != 110 {
		t.Error("la possition n'est pas celle qui est prédéfini par SpawnX et SpawnY")
	}
}

func TestAddParticuleRandomSpeed(t *testing.T) {
	resetConfigAddParticule()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.SpeedType = 1

	sys.AddParticule()
	sys.AddParticule()

	if sys.Content.Front().Value.(*particles.Particle).SpeedX == sys.Content.Front().Next().Value.(*particles.Particle).SpeedX && sys.Content.Front().Value.(*particles.Particle).SpeedY == sys.Content.Front().Next().Value.(*particles.Particle).SpeedY {
		t.Error("les vitesse ne sont pas aléatoire alors que TypeSpeed n'est pas égale à 0")
	}
}

func TestAddParticuleSeedObject(t *testing.T) {
	resetConfigAddParticule()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.SpeedType = 1
	config.General.SpawnOnAnObject = true
	config.General.SpawnObject = "circle"
	config.General.SpawnObjectWidth = 500

	sys.AddParticule()
	sys.AddParticule()

	if sys.Content.Front().Value.(*particles.Particle).SpeedX == sys.Content.Front().Next().Value.(*particles.Particle).SpeedX && sys.Content.Front().Value.(*particles.Particle).SpeedY == sys.Content.Front().Next().Value.(*particles.Particle).SpeedY {
		t.Error("les vitesse ne sont pas aléatoire alors que TypeSpeed n'est pas égale à 0")
	}
}

func TestAddParticuleSpeed0(t *testing.T) {
	resetConfigAddParticule()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.SpeedType = 0

	sys.AddParticule()

	if sys.Content.Front().Value.(*particles.Particle).SpeedX != 0 || sys.Content.Front().Value.(*particles.Particle).SpeedY != 0 {
		t.Error("la vitesse n'est 0 alors que TypeSpeed est égale à 0")
	}
}

func TestAddParticuleConstantLife(t *testing.T) {
	resetConfigAddParticule()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.Life = 10

	sys.AddParticule()

	if sys.Content.Front().Value.(*particles.Particle).Life != 10 || sys.Content.Front().Value.(*particles.Particle).LifeInit != 10 {
		t.Error("la vie ou la vie initiale n'est pas à 10")
	}
}

func TestAddParticuleRandomLife(t *testing.T) {
	resetConfigAddParticule()
	sys := System{Content: list.New(), DeadList: list.New()}

	config.General.HaveLife = true
	config.General.RandomLife = true
	config.General.Life = 10

	sys.AddParticule()
	sys.AddParticule()

	if sys.Content.Front().Value.(*particles.Particle).Life == sys.Content.Front().Value.(*particles.Particle).Life || sys.Content.Front().Value.(*particles.Particle).LifeInit != sys.Content.Front().Value.(*particles.Particle).Life || sys.Content.Front().Next().Value.(*particles.Particle).LifeInit != sys.Content.Front().Value.(*particles.Particle).Life {
		t.Error("les vies initiale ne sont pas à 50 ou elles ne sont pas aléatoire")
	}
}

//
//
//
//
//
//

func resetConfigAddParticule() {
	{
		config.General.WindowSizeX = 100
		config.General.WindowSizeY = 100
		config.General.RandomSpawn = false
		config.General.SpawnCenter = false
		config.General.SpawnX = 0
		config.General.SpawnY = 0
		config.General.SpawnOnAnObject = false
		config.General.SpawnObject = ""
		config.General.SpawnObjectWidth = 0
		config.General.Rotation = 0
		config.General.ScaleX = 1
		config.General.ScaleY = 1
		config.General.Opacity = 1
		config.General.ColorRed = 1
		config.General.ColorGreen = 1
		config.General.ColorBlue = 1
		config.General.SpeedType = 1
		config.General.HaveLife = false
		config.General.RandomLife = false
		config.General.Life = 50
		config.General.Collision = false
		config.General.CollisionAmongParticle = false

	}
}
