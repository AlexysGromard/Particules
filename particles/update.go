package particles

import (
	//"math/rand"
	"project-particles/config"
	//"time"
	//"fmt"
)

var part_particle float64

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	part_particle += config.General.SpawnRate

	//actualisation des particules
	for i := s.Content.Front(); i != nil; i = i.Next() {
		p, ok := i.Value.(*Particle)

		if !((p.PositionX < 0-config.General.MarginOutsideScreen || p.PositionX > float64(config.General.WindowSizeX) + config.General.MarginOutsideScreen) || (p.PositionY < 0-config.General.MarginOutsideScreen || p.PositionY > float64(config.General.WindowSizeY) + config.General.MarginOutsideScreen) || p.Life <= 0 ){
			if ok {
				// mouvement des particule
				p.SpeedY = p.SpeedY + config.General.Gravity

				p.PositionX = p.PositionX + p.SpeedX
				p.PositionY = p.PositionY + p.SpeedY
		

				if config.General.HaveLife{
					p.Life -= 1
				}

				/*
				//suprestion des particule sortante
				if config.General.KillParticlesOutside{
					if (p.PositionX+p.ScaleX <= 0 && p.PositionX-p.ScaleX >= float64(config.General.WindowSizeX)) && (p.PositionY+p.ScaleY <= 0 && p.PositionX-p.ScaleY >= float64(config.General.WindowSizeY)){
						s.Content.Remove(i)
					}
				}
				*/

				UpdateColor(p)
				UpdateScale(p)
				UpdateRotation(p)
				UpdateOpacity(p)


			}
		}
	}

	//ajout des particule
	if part_particle >=1{
		
		for i := 0; float64(i) < part_particle; i++ {
			part_particle -= 1
			Creat_Particles(s.Content)
		}
	}

}
