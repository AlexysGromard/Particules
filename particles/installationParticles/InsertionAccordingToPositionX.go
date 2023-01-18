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
			for NuméroParticule := list.Front(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Next() {
				pList, _ := NuméroParticule.Value.(*particles.Particle)
				if p.PositionX < pList.PositionX {
					list.InsertBefore(p, NuméroParticule)
					break
				} else if NuméroParticule.Next() == nil {
					list.PushBack(p)
					break
				}
			}
		} else {
			for NuméroParticule := list.Back(); NuméroParticule != nil; NuméroParticule = NuméroParticule.Prev() {
				pList, _ := NuméroParticule.Value.(*particles.Particle)
				if p.PositionX > pList.PositionX {
					list.InsertAfter(p, NuméroParticule)
					break
				} else if NuméroParticule.Next() == nil {
					list.PushFront(p)
					break
				}
			}
		}
	} else {
		list.PushFront(p)
	}

}
