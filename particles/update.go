package particles

//import "../project-particles/particles"

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	for i:=s.Content.Front(); i != nil; i = i.Next(){

		p, ok := i.Value.(*Particle)
		if ok{
			p.PositionX = p.PositionX + p.SpeedX
			p.PositionY = p.PositionY + p.SpeedY
		}
	}
}
