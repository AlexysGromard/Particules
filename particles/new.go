package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	rand.Seed(time.Now().UnixNano())
	// Initialisation du System
	Particles := System{Content: list.New()}
	// Initialisation des particules
	for i := 0; i < config.General.InitNumParticles; i++ {
		Particles.Add_Particule()
	}
	return Particles
}
