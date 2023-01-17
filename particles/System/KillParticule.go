package System

import (
	"container/list"
	"project-particles/particles"
)


func (s *System) KillParticule(e *list.Element) {
	p, _ := e.Value.(*particles.Particle)
	s.DeadList.PushFront(p)
	s.Content.Remove(e)
}
