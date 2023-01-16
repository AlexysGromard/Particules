package particles

import (
	"container/list"
)

func (s *System) PlaceAccordingToPosition() {
	i := s.Content.Front()
	for i != nil{
		suivant := i.Next()

		Devant   := i.Next()
		Derrière := i.Prev()		
		rec(s,i,Derrière,Devant)

		i = suivant
	}
}

func rec(s *System,e,Derrière, Devant *list.Element){
	p, _ := e.Value.(*Particle)
	if Derrière == nil && Devant != nil{
		pDevant,_ := Devant.Value.(*Particle)
		if p.PositionX > pDevant.PositionX{
			rec(s,e,Derrière,Devant.Next())
		}else{
			s.Content.MoveBefore(e, Devant)
		}
	}else if Derrière != nil && Devant == nil{
		pDerrière,_ := Derrière.Value.(*Particle)
		if p.PositionX < pDerrière.PositionX{
			rec(s,e,Derrière.Prev(),Derrière)
		}else{
			s.Content.MoveAfter(e, Derrière)
		}
	}else if Derrière != nil && Devant != nil{
		pDevant,_ := Devant.Value.(*Particle)
		pDerrière,_ := Derrière.Value.(*Particle)

		if p.PositionX > pDevant.PositionX{
			rec(s,e,Devant,Devant.Next())
		}else if p.PositionX < pDerrière.PositionX{
			rec(s,e,Derrière.Prev(),Derrière)
		}else{
			s.Content.MoveBefore(e, Devant)
		}

	}
}
