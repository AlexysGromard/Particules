package particles

import (
	"project-particles/config"
)

// cette fonction crée des particules en fonction des paramètre indiqué dans "config.json"
// si DeadListe contient des particules alors il réutilise les particule comptenu dans DeadList

//enter : DeadList (une liste de particule morte qui sont aussi dans le System)
func (s *System) Add_Particule() {

	// Initialisation de la position de la particule
	var PositionX float64
	var PositionY float64
	if config.General.RandomSpawn{
		PositionX, PositionY = Random_Position(config.General.WindowSizeX, config.General.WindowSizeY)
	} else {
		PositionX = float64(config.General.SpawnX)
		PositionY = float64(config.General.SpawnY)
	}

	// Initialisation de la Vitesse de la partucle
	var SpeedX, SpeedY float64
	SpeedX, SpeedY = Random_Speed(config.General.SpeedType)
	

	//Ajoute de la paticule au System
	s.Content.PushFront(&Particle{
		PositionX: PositionX,
		PositionY: PositionY,
		Rotation:  config.General.Rotation,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
		ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
		Opacity:  config.General.Opacity,
		SpeedX:   SpeedX,
		SpeedY:   SpeedY,
	})
	
}
