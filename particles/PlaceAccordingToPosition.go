package particles

import (
	"container/list"
)

func (s *System) PlaceAccordingToPosition(e *list.Element) {
	p, ok := e.Value.(*Particle)
	//fmt.Println(" 1 : ", p.ColorRed, "/", p.ColorGreen, "/", p.ColorBlue)
	if ok {
		if p.SpeedX > 0 {
			eNext := e.Next()
			for eNext != nil {
				pNext, ok := eNext.Value.(*Particle)
				if ok && pNext.PositionX > p.PositionX {
					s.Content.MoveBefore(e, eNext)
					break
				}
				eNext = eNext.Next()
			}

		} else if p.SpeedX < 0 {
			ePrev := e.Prev()
			for ePrev != nil {
				pPrev, ok := ePrev.Value.(*Particle)
				if ok {
					if pPrev.PositionX > p.PositionX {
						s.Content.MoveAfter(e, ePrev)
						break
					}
				}
				ePrev = ePrev.Prev()
			}
		}
	}

}
