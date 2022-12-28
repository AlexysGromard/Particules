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

	var _Life int

	if config.General.HaveLife && config.General.RandomLife{
		_Life = rand.Intn(180)
	}else{
		_Life = config.General.Life
	}

	l.PushFront(&Particle{
		PositionX: _PositionX,
		PositionY: _PositionY,
		Rotation: config.General.Rotation,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
		ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
		Opacity: config.General.Opacity,
		SpeedX:  (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1),
		SpeedY:  (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1),
		Life : _Life,
	})

}

func UpdateColor(p *Particle){
	if config.General.ChangeColorAccordingTo == 1{
		p.ColorRed = ((p.PositionX + p.PositionY)/float64(config.General.WindowSizeX + config.General.WindowSizeY))*config.General.ColorRed
		p.ColorGreen = ((p.PositionX + p.PositionY)/float64(config.General.WindowSizeX + config.General.WindowSizeY))*config.General.ColorGreen
		p.ColorBlue = ((p.PositionX + p.PositionY)/float64(config.General.WindowSizeX + config.General.WindowSizeY))*config.General.ColorBlue
	}else if config.General.ChangeColorAccordingTo == 2{
		p.ColorRed = (float64(p.Life)/float64(config.General.Life))*config.General.ColorRed
		p.ColorGreen = (float64(p.Life)/float64(config.General.Life))*config.General.ColorGreen
		p.ColorBlue = (float64(p.Life)/float64(config.General.Life))*config.General.ColorBlue
	}
}

func UpdateScale(p *Particle){
	if config.General.ChangeScaleAccordingTo == 1{
		p.ScaleX = ((p.PositionX + p.PositionY)/float64(config.General.WindowSizeX + config.General.WindowSizeY))*config.General.ScaleX
		p.ScaleY = ((p.PositionX + p.PositionY)/float64(config.General.WindowSizeX + config.General.WindowSizeY))*config.General.ScaleY
	}else if config.General.ChangeScaleAccordingTo == 2{
		p.ScaleX = (float64(p.Life)/float64(config.General.Life))*config.General.ScaleX
		p.ScaleY = (float64(p.Life)/float64(config.General.Life))*config.General.ScaleY
	}
}

func UpdateRotation(p *Particle){
	if config.General.ChangeRotationAccordingTo == 1{
		p.Rotation = ((p.PositionX + p.PositionY)/float64(config.General.WindowSizeX + config.General.WindowSizeY))*config.General.Rotation
	}else if config.General.ChangeRotationAccordingTo == 2{
		p.Rotation = (float64(p.Life)/float64(config.General.Life))*config.General.Rotation
	}
}

func UpdateOpacity(p *Particle){
	if config.General.ChangeOpacityAccordingTo == 1{
		p.Opacity = ((p.PositionX + p.PositionY)/float64(config.General.WindowSizeX + config.General.WindowSizeY))*config.General.Opacity
	}else if config.General.ChangeOpacityAccordingTo == 2{
		p.Opacity = (float64(p.Life)/float64(config.General.Life))*config.General.Opacity
	}
}

