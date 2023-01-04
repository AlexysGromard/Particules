package particles

import (
	"container/list"
	"project-particles/config"
)

// cette fonction crée des particules en fonction des paramètre indiqué dans "config.json"
func (s *System) Add_Particule(DeadList *list.List) {

	var CentreX, CentreY float64
	CentreX, CentreY = float64(config.General.WindowSizeX)/2, float64(config.General.WindowSizeY)/2

	// Initialisation de la position de la particule
	var PositionX float64
	var PositionY float64
	if config.General.RandomSpawn && !config.General.SpawnOnAnObject {
		PositionX, PositionY = Random_Position(config.General.WindowSizeX, config.General.WindowSizeY)
	} else if config.General.SpawnOnAnObject {
		PositionX, PositionY = PositionAccordingToShape(config.General.SpawnObject, config.General.SpawnObjectWidth, CentreX, CentreY)
	} else {
		PositionX = float64(config.General.SpawnX)
		PositionY = float64(config.General.SpawnY)
	}

	var SpeedX, SpeedY float64
	if config.General.SpawnOnAnObject {
		SpeedX, SpeedY = SpeedAccordingToShape(config.General.SpeedType, PositionX, PositionY, CentreX, CentreY)
	} else {
		SpeedX, SpeedY = Random_Speed(config.General.SpeedType)
	}

	// Initialisation de la vie de la particule
	var Life int
	if config.General.HaveLife && config.General.RandomLife {
		Life = Random_Life()
	} else {
		Life = config.General.Life
	}

	ParticuleDead := DeadList.Front()
	if ParticuleDead != nil {
		pd, ok := ParticuleDead.Value.(*Particle)
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
			pd.Life = Life

			DeadList.Remove(ParticuleDead)
		}

	} else {
		s.Content.PushFront(&Particle{
			PositionX: PositionX,
			PositionY: PositionY,
			Rotation:  config.General.Rotation,
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
			ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
			Opacity: config.General.Opacity,
			SpeedX:  SpeedX,
			SpeedY:  SpeedY,
			Life:    Life,
		})
	}
}
