package system

import (
	"project-particles/config"
	"project-particles/particles"
	"project-particles/particles/installationParticles"
)

// cette fonction crée des particules en fonction des paramètre indiqué dans "config.json"
// si DeadListe contient des particules alors il réutilise les particule comptenu dans DeadList

// enter : DeadList (une liste de particule morte qui sont aussi dans le System)
func (s *System) addParticule() {

	var CentreX, CentreY float64
	CentreX, CentreY = float64(config.General.WindowSizeX)/2, float64(config.General.WindowSizeY)/2

	// Initialisation de la position de la particule
	var PositionX float64
	var PositionY float64
	if config.General.RandomSpawn && !config.General.SpawnOnAnObject {
		PositionX, PositionY = installationParticles.Random_Position(config.General.WindowSizeX, config.General.WindowSizeY)
	} else if config.General.SpawnOnAnObject {
		PositionX, PositionY = installationParticles.PositionAccordingToShape(config.General.SpawnObject, config.General.SpawnObjectWidth, CentreX, CentreY)
	} else {
		PositionX = float64(config.General.SpawnX)
		PositionY = float64(config.General.SpawnY)
	}

	// Initialisation de la Vitesse de la partucle
	var SpeedX, SpeedY float64
	if config.General.SpawnOnAnObject {
		SpeedX, SpeedY = installationParticles.SpeedAccordingToShape(config.General.SpeedType, PositionX, PositionY, CentreX, CentreY)
	} else if config.General.SpeedType == 0 {
		SpeedX, SpeedY = 0, 0
	} else {
		SpeedX, SpeedY = installationParticles.Random_Speed(config.General.SpeedType)
	}

	// Initialisation de la vie de la particule
	var Life int
	if config.General.HaveLife && config.General.RandomLife {
		Life = installationParticles.Random_Life(50)
	} else {
		Life = config.General.Life
	}

	//Ajoute de la paticule au System
	ParticuleDead := s.DeadList.Front()
	if ParticuleDead != nil {
		pd, ok := ParticuleDead.Value.(*particles.Particle)
		if ok {
			pd.PositionX = PositionX
			pd.PositionY = PositionY
			pd.Rotation = config.General.Rotation
			pd.ScaleX = config.General.ScaleX
			pd.ScaleY = config.General.ScaleY
			pd.ColorRed = config.General.ColorRed
			pd.ColorGreen = config.General.ColorGreen
			pd.ColorBlue = config.General.ColorBlue
			pd.Opacity = config.General.Opacity
			pd.SpeedX = SpeedX
			pd.SpeedY = SpeedY
			pd.LifeInit = Life
			pd.Life = Life

			s.resurrectParticule(ParticuleDead, float64(config.General.WindowSizeX)/2, config.General.Collision, config.General.CollisionAmongParticle)
		}

	} else {
		NouvelleParticule := &particles.Particle{
			PositionX: PositionX,
			PositionY: PositionY,
			Rotation:  config.General.Rotation,
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
			ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
			Opacity:  config.General.Opacity,
			SpeedX:   SpeedX,
			SpeedY:   SpeedY,
			LifeInit: Life, Life: Life,
		}

		if config.General.Collision && config.General.CollisionAmongParticle {
			installationParticles.InsertionAccordingToPositionX(s.Content, NouvelleParticule, float64(config.General.WindowSizeX)/2)
		} else {
			s.Content.PushFront(NouvelleParticule)
		}
	}
}
