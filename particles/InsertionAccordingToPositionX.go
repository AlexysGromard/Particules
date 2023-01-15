package particles

func (s *System) InsertionAccordingToPositionX(p *Particle, Xmoitier float64, Collision, CollisionAmongParticle bool) (ok bool) {
	if Collision && CollisionAmongParticle && s.Content.Len() != 0 {
		if p.PositionX < Xmoitier {
			for i := s.Content.Front(); i != nil; i = i.Next() {
				pList, _ := i.Value.(*Particle)
				if p.PositionX < pList.PositionX {
					s.Content.InsertBefore(p, i)
					return true
				} else if i.Next() == nil {
					s.Content.PushBack(p)
					return true
				}
			}
		} else {
			for i := s.Content.Back(); i != nil; i = i.Prev() {
				pList, _ := i.Value.(*Particle)
				if p.PositionX > pList.PositionX {
					s.Content.InsertAfter(p, i)
					return true
				} else if i.Next() == nil {
					s.Content.PushFront(p)
					return true
				}
			}
		}

		return false
		// la fonction renvois faux car elle n'est pas rentrer dans la boucle for ce qui veux dire
		//que la liste s.content est vide et donc qu'on ne peux pas ajouter de particule en fonction
		//de la position
	} else {
		s.Content.PushFront(p)
		return true
	}

}
