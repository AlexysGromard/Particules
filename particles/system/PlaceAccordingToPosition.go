package system

import (
	"container/list"
	"project-particles/particles"
)

// La fonction PlaceAccordingToPosition permet de déplacer les particules dans la liste
// en fonction de leur position en X
func (s *System) placeAccordingToPosition() {
	i := s.Content.Front()
	for i != nil {
		suivant := i.Next()

		Devant := i.Next()
		Derrière := i.Prev()
		location(s, i, Derrière, Devant)

		i = suivant
	}
}

// La fonction location permet de déplacer les particules dans la liste
// en fonction de leur position en X
// Elle prend en paramètre un élément de la liste des particules
func location(s *System, e, Derrière, Devant *list.Element) {
	p, _ := e.Value.(*particles.Particle)
	if Derrière == nil && Devant != nil { // Si on est au début de la liste
		pDevant, _ := Devant.Value.(*particles.Particle)
		if p.PositionX > pDevant.PositionX {
			location(s, e, Derrière, Devant.Next())
		} else {
			s.Content.MoveBefore(e, Devant)
		}
	} else if Derrière != nil && Devant == nil { // Si on est à la fin de la liste
		pDerrière, _ := Derrière.Value.(*particles.Particle)
		if p.PositionX < pDerrière.PositionX {
			location(s, e, Derrière.Prev(), Derrière)
		} else {
			s.Content.MoveAfter(e, Derrière)
		}
	} else if Derrière != nil && Devant != nil { // Si on est au milieu de la liste
		pDevant, _ := Devant.Value.(*particles.Particle)
		pDerrière, _ := Derrière.Value.(*particles.Particle)
		if p.PositionX > pDevant.PositionX {
			location(s, e, Devant, Devant.Next())
		} else if p.PositionX < pDerrière.PositionX {
			location(s, e, Derrière.Prev(), Derrière)
		} else {
			s.Content.MoveBefore(e, Devant)
		}
	}
}
