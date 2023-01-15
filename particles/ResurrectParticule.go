package particles

import "container/list"

func (s *System) ResurrectParticule(e *list.Element, Xmoitier float64, Collision, CollisionAmongParticle bool) {
	p, _ := e.Value.(*Particle)

	s.InsertionAccordingToPositionX(p, Xmoitier, Collision, CollisionAmongParticle)
	s.DeadList.Remove(e)
}
