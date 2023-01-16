package particles

func (s *System) InsertionAccordingToPositionX(p *Particle, Xmoitier float64) {
		if s.Content.Len() != 0{
			
			if p.PositionX < Xmoitier {

				for i := s.Content.Front(); i != nil; i = i.Next() {
					pList, _ := i.Value.(*Particle)
					if p.PositionX < pList.PositionX {
						s.Content.InsertBefore(p, i)
						break
					} else if i.Next() == nil {
						s.Content.PushBack(p)
						break
					}
				}
			} else {
				for i := s.Content.Back(); i != nil; i = i.Prev() {
					pList, _ := i.Value.(*Particle)
					if p.PositionX > pList.PositionX {
						s.Content.InsertAfter(p, i)
						break
					} else if i.Next() == nil {
						s.Content.PushFront(p)
						break
					}
				}
			}
		}else{
			s.Content.PushFront(p)
		}

		// la fonction renvois faux car elle n'est pas rentrer dans la boucle for ce qui veux dire
		//que la liste s.content est vide et donc qu'on ne peux pas ajouter de particule en fonction
		//de la position
	

}
