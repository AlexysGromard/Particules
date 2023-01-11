package particles

import "container/list"

func (s *System) ResurrectParticule(e *list.Element) {
	p, _ := e.Value.(*Particle)
	s.Content.PushFront(p)
	s.DeadList.Remove(e)
}
