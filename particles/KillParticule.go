package particles

import (
	"container/list"
)

func (s *System) KillParticule(e *list.Element) {
	p, _ := e.Value.(*Particle)
	s.DeadList.PushFront(p)
	s.Content.Remove(e)
}
