package system

import (
	"container/list"
	"project-particles/particles"
	"project-particles/particles/installationParticles"
)

func (s *System) ResurrectParticule(e *list.Element, Xmoitier float64, Collision, CollisionAmongParticle bool) {
	p, _ := e.Value.(*particles.Particle)
	if Collision && CollisionAmongParticle {
		installationParticles.InsertionAccordingToPositionX(s.Content, p, Xmoitier)

	} else {
		s.Content.PushFront(p)
	}
	s.DeadList.Remove(e)
}
