package system

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
	// Initialisation de la génération aléatoire
	rand.Seed(time.Now().UnixNano())
	// Initialisation de la liste des particules
	system := System{Content: list.New(), DeadList: list.New()}
	// Initialisation des particules
	for NuméroToure := 0; NuméroToure < config.General.InitNumParticles; NuméroToure++ {
		// Initialisation de la position de la particule
		system.addParticule()
	}
	return system
}
