package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	//"time"
)

//cette fonction crée des particules en fonction des paramètre indiqué dans "config.json"
func Creat_Particles(l *list.List){

	var _PositionX float64
	var _PositionY float64

	// Initialisation de la position de la particule
	if config.General.RandomSpawn {
		_PositionX = float64(rand.Intn(config.General.WindowSizeX)) + rand.Float64()
		_PositionY = float64(rand.Intn(config.General.WindowSizeY)) + rand.Float64()
	} else {
		_PositionX = float64(config.General.SpawnX)
		_PositionY = float64(config.General.SpawnY)
	}
	// Initialisation de la vitesse de la particule
	var minSpeed, maxSpeed int
	switch config.General.SpeedType {
	case 1:
		minSpeed = 0
		maxSpeed = 1
	case 2:
		minSpeed = 1
		maxSpeed = 3
	case 3:
		minSpeed = 3
		maxSpeed = 5
	}

	l.PushFront(&Particle{
		PositionX: _PositionX,
		PositionY: _PositionY,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
		ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
		Opacity: config.General.Opacity,
		SpeedX:  (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1),
		SpeedY:  (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1),
	})

}