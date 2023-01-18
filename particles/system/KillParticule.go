package system

import (
	"container/list"
	"project-particles/particles"
)

// La fonction KillParticule permet de supprimer une particule de la liste des particules
// Elle prend en paramètre un élément de la liste des particules
// Elle ajoute la particule à la liste des particules mortes
// Elle supprime l'élément de la liste des particules
func (s *System) killParticule(e *list.Element) {
	p, _ := e.Value.(*particles.Particle)
	s.DeadList.PushFront(p)
	s.Content.Remove(e)
}
