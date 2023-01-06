package particles

import (
	"project-particles/config"
)

var Part_particle float64 // cette variable correspond au partie de particule

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	Part_particle += config.General.SpawnRate

	//actualisation des particules
	for i := s.Content.Front(); i != nil; i = i.Next() {
		p, ok := i.Value.(*Particle)

		if ok {
			// mouvement des particule
			//la position est modifie ne fonction de la vitesse
			p.PositionX = p.PositionX + p.SpeedX
			p.PositionY = p.PositionY + p.SpeedY
			//
		}
	}
	//

	//ajout des particule
	if Part_particle >= 1 {

		for i := 0; float64(i) < Part_particle; i++ {
			Part_particle -= 1
			s.Add_Particule()
		}
	}
	//
}
