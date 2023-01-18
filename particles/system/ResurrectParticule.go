package system

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/installationParticles"
)

// La méthode ResurrectParticule permet de réanimer une particule
// Elle prend en paramètre un élément de la liste des particules mortes
// Elle ajoute la particule à la liste des particules
// Elle supprime l'élément de la liste des particules mortes
func (s *System) resurrectParticule(e *list.Element, Xmoitier float64, Collision, CollisionAmongParticle bool) {
	p, _ := e.Value.(*particles.Particle)
	if Collision && CollisionAmongParticle {
		installationParticles.InsertionAccordingToPositionX(s.Content, p, Xmoitier)

	} else {
		s.Content.PushFront(p)
	}
	s.DeadList.Remove(e)
}
