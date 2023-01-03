package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"time"
)

//cette variable est un System qui contiendra les particule morte
var DeadParticles System

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	//l := list.New()

	/* //créer une particule //
		l.PushFront(&Particle{
		PositionX: float64(config.General.WindowSizeX) / 2,
		PositionY: float64(config.General.WindowSizeY) / 2,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
	})
	*/

	rand.Seed(time.Now().UnixNano())

	//
	DeadParticles = System{Content: list.New()}
	Particles := System{Content: list.New()}
	//

	for i := 0; i < config.General.InitNumParticles; i++ {
		// Initialisation de la position de la particule

		Particles.Add_Particule()

	}

	return Particles
}
