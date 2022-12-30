package particles

import (
	"container/list"
	"math"
	"math/rand"
	"project-particles/config"
	//"time"
	//"fmt"
)

// cette fonction crée des particules en fonction des paramètre indiqué dans "config.json"
func Creat_Particles(l *list.List) {

	// Initialisation de la position de la particule
	var _PositionX float64
	var _PositionY float64
	if config.General.RandomSpawn && !config.General.SpawnOnAnObject {
		_PositionX = float64(rand.Intn(config.General.WindowSizeX)) + rand.Float64()
		_PositionY = float64(rand.Intn(config.General.WindowSizeY)) + rand.Float64()
	} else if config.General.SpawnOnAnObject {
		var xCenter float64 = float64(config.General.WindowSizeX / 2)
		var yCenter float64 = float64(config.General.WindowSizeY / 2)
		switch config.General.SpawnObject {
		case "circle":
			var randomPosition int = rand.Intn(360)
			_PositionX = xCenter + float64(float64(config.General.SpawnObjectWidth/2)*math.Cos(float64(randomPosition)))
			_PositionY = yCenter + float64(float64(config.General.SpawnObjectWidth/2)*math.Sin(float64(randomPosition)))
		case "square":
			var randomPosition int = rand.Intn(4)
			switch randomPosition {
			case 0:
				_PositionX = xCenter + float64(config.General.SpawnObjectWidth/2)
				_PositionY = yCenter + float64(rand.Intn(config.General.SpawnObjectWidth)) - float64(config.General.SpawnObjectWidth/2)
			case 1:
				_PositionX = xCenter + float64(rand.Intn(config.General.SpawnObjectWidth)) - float64(config.General.SpawnObjectWidth/2)
				_PositionY = yCenter + float64(config.General.SpawnObjectWidth/2)
			case 2:
				_PositionX = xCenter - float64(config.General.SpawnObjectWidth/2)
				_PositionY = yCenter + float64(rand.Intn(config.General.SpawnObjectWidth)) - float64(config.General.SpawnObjectWidth/2)
			case 3:
				_PositionX = xCenter + float64(rand.Intn(config.General.SpawnObjectWidth)) - float64(config.General.SpawnObjectWidth/2)
				_PositionY = yCenter - float64(config.General.SpawnObjectWidth/2)
			}
		}
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

	// Initialisation de la vie de la particule
	var _Life int
	if config.General.HaveLife && config.General.RandomLife {
		_Life = rand.Intn(180)
	} else {
		_Life = config.General.Life
	}

	ParticuleDead := DeadParticles.Content.Front()

	if ParticuleDead != nil {
		pd, ok := ParticuleDead.Value.(*Particle)

		if ok {
			pd.PositionX = _PositionX
			pd.PositionY = _PositionY
			pd.Rotation = config.General.Rotation
			pd.ScaleX = config.General.ScaleX
			pd.ScaleY = config.General.ScaleY
			pd.ColorRed = config.General.ColorRed
			pd.ColorGreen = config.General.ColorGreen
			pd.ColorBlue = config.General.ColorBlue
			pd.Opacity = config.General.Opacity
			pd.SpeedX = (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1)
			pd.SpeedY = (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1)
			pd.Life = _Life

			DeadParticles.Content.Remove(ParticuleDead)
		}

	} else {
		l.PushFront(&Particle{
			PositionX: _PositionX,
			PositionY: _PositionY,
			Rotation:  config.General.Rotation,
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
			ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
			Opacity: config.General.Opacity,
			SpeedX:  (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1),
			SpeedY:  (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1),
			Life:    _Life,
		})
	}

	//fmt.Println(DeadParticles.Content)
}

// Cette fonction actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeColorAccordingTo" dans config.json
func UpdateColor(p *Particle) {
	if config.General.ChangeColorAccordingTo == 1 {
		p.ColorRed = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorRed
		p.ColorGreen = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorGreen
		p.ColorBlue = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ColorBlue
	} else if config.General.ChangeColorAccordingTo == 2 {
		p.ColorRed = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorRed
		p.ColorGreen = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorGreen
		p.ColorBlue = (float64(p.Life) / float64(config.General.Life)) * config.General.ColorBlue
	}
}

// Cette fonction actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeScaleAccordingTo" dans config.json
func UpdateScale(p *Particle) {
	if config.General.ChangeScaleAccordingTo == 1 {
		p.ScaleX = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ScaleX
		p.ScaleY = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.ScaleY
	} else if config.General.ChangeScaleAccordingTo == 2 {
		p.ScaleX = (float64(p.Life) / float64(config.General.Life)) * config.General.ScaleX
		p.ScaleY = (float64(p.Life) / float64(config.General.Life)) * config.General.ScaleY
	}
}

// Cette fonction actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeRotationAccordingTo" dans config.json
func UpdateRotation(p *Particle) {
	if config.General.ChangeRotationAccordingTo == 1 {
		p.Rotation = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.Rotation
	} else if config.General.ChangeRotationAccordingTo == 2 {
		p.Rotation = (float64(p.Life) / float64(config.General.Life)) * config.General.Rotation
	}
}

// Cette fonction actualise la couleur de la particule "p" en fonction du nombre choisie pour "ChangeOpacityAccordingTo" dans config.json
func UpdateOpacity(p *Particle) {
	if config.General.ChangeOpacityAccordingTo == 1 {
		p.Opacity = ((p.PositionX + p.PositionY) / float64(config.General.WindowSizeX+config.General.WindowSizeY)) * config.General.Opacity
	} else if config.General.ChangeOpacityAccordingTo == 2 {
		p.Opacity = (float64(p.Life) / float64(config.General.Life)) * config.General.Opacity
	}
}
