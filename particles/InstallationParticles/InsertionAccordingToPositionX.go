package InstallationParticles

import 	(
	"container/list"
	"project-particles/particles"
)

func InsertionAccordingToPositionX(list *list.List,p *particles.Particle, Xmoitier float64) {
		if list.Len() != 0{
			
			if p.PositionX < Xmoitier {
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
		}else{
			list.PushFront(p)
		}

}
