package particles

import (
	"math/rand"
	"project-particles/config"
	"time"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	for i := s.Content.Front(); i != nil; i = i.Next() {
		p, ok := i.Value.(*Particle)
		if ok {
			p.PositionX = p.PositionX + p.SpeedX
			p.PositionY = p.PositionY + p.SpeedY
		}
	}

	var _PositionX float64
	var _PositionY float64
	rand.Seed(time.Now().UnixNano())
	for i := 0; float64(i) < config.General.SpawnRate; i++ {

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
		s.Content.PushFront(&Particle{
			PositionX: _PositionX,
			PositionY: _PositionY,
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
			ColorRed: config.General.ColorRed, ColorGreen: config.General.ColorGreen, ColorBlue: config.General.ColorBlue,
			Opacity: config.General.Opacity,
			SpeedX:  (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1),
			SpeedY:  (float64(rand.Intn(maxSpeed-minSpeed)+minSpeed) + rand.Float64()) * float64(rand.Intn(2)*2-1),
		})
	}

}
