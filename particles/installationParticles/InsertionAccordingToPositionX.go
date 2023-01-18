package installationParticles

import (
	"container/list"
	"project-particles/particles"
)

// La fonction InsertionAccordingToPositionX ajoute une particule à la liste selon sa position X
// Si la liste est vide, la particule est ajouté à la liste
// Si la particule est plus petite que la moitié de la liste, la particule est ajouté à la liste en partant du début
// Si la particule est plus grande que la moitié de la liste, la particule est ajouté à la liste en partant de la fin
func InsertionAccordingToPositionX(list *list.List, p *particles.Particle, Xhalf float64) {
	if list.Len() != 0 {

		if p.PositionX < Xhalf {
			for i := list.Front(); i != nil; i = i.Next() {
				pList, _ := i.Value.(*particles.Particle)
				if p.PositionX < pList.PositionX {
					list.InsertBefore(p, i)
					break
				} else if i.Next() == nil {
					list.PushBack(p)
					break
				}
			}
		} else {
			for i := list.Back(); i != nil; i = i.Prev() {
				pList, _ := i.Value.(*particles.Particle)
				if p.PositionX > pList.PositionX {
					list.InsertAfter(p, i)
					break
				} else if i.Next() == nil {
					list.PushFront(p)
					break
				}
			}
		}
	} else {
		list.PushFront(p)
	}

}
