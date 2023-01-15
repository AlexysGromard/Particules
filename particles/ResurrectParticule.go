package particles

import "container/list"

func (s *System) ResurrectParticule(e *list.Element,CollisionAmongParticle bool) {
	p, _ := e.Value.(*Particle)

	s.InsertionAccordingToPositionX(p,CollisionAmongParticle)
	/*
	for i := s.Content.Front();i != nil ; i = i.Next(){
		pList,_ := i.Value.(*Particle)
		if p.PositionX < pList.PositionX{
			s.Content.InsertBefore(p,i)
			break
		}
	}*/
	s.DeadList.Remove(e)
}
