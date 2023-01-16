package particles

import (
	"container/list"

)

func (s *System) ResurrectParticule(e *list.Element, Xmoitier float64, Collision, CollisionAmongParticle bool) {
	p, _ := e.Value.(*Particle)
	if Collision && CollisionAmongParticle{
		s.InsertionAccordingToPositionX(p, Xmoitier)
		
	}else{
		s.Content.PushFront(p)
	}
	s.DeadList.Remove(e)
}
	
