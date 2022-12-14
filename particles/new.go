package particles

import (
	"container/list"
	"project-particles/config"
	"math/rand"
	"time"
	"fmt"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()

	/* //créer une particule //
				l.PushFront(&Particle{
				PositionX: float64(config.General.WindowSizeX) / 2,
				PositionY: float64(config.General.WindowSizeY) / 2,
				ScaleX:    1, ScaleY: 1,
				ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
				Opacity: 1,
			}) 
	*/

	var _PositionX float64
	var _PositionY float64

	rand.Seed(time.Now().UnixNano())

	for i:=0;i < config.General.InitNumParticles ; i++ {
		// Initialisation de la position de la particule
		if config.General.RandomSpawn{
			_PositionX = float64(rand.Intn(config.General.WindowSizeX)) + rand.Float64()
			_PositionY = float64(rand.Intn(config.General.WindowSizeY)) + rand.Float64()
		}else{
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
			SpeedX : float64(rand.Intn(maxSpeed*2-minSpeed)+minSpeed) + rand.Float64() - 10,
			SpeedY : float64(rand.Intn(maxSpeed*2-minSpeed)+minSpeed) + rand.Float64() - 10,
		})
	}

	

	
	return System{Content: l}
}
