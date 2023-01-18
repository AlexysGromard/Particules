package system

import (
	"project-particles/config"
	"project-particles/particles"
	"project-particles/particles/ParticleModification"
)

var Part_particle float64 // cette variable correspond au partie de particule

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	Part_particle += config.General.SpawnRate
	particleNumber := s.Content.Front()
	//actualisation des particules
	for particleNumber != nil {

		p, ready := particleNumber.Value.(*particles.Particle)

		particle := particleNumber.Next()

		if ready {

			// mouvement des particule
			//la vitesse est modifier en fonction de l'accélération de la gravité
			p.SpeedY = p.SpeedY + config.General.Gravity
			//

			//la position est modifie ne fonction de la vitesse
			p.PositionX = p.PositionX + p.SpeedX
			p.PositionY = p.PositionY + p.SpeedY
			//
			//

			//quand le system de vie est active on vas retirer une uniter de vis à chaque update
			if config.General.HaveLife {
				p.Life -= 1
			}
			//

			//tue les particules morte
			if (ParticleModification.ParticleOutsideXLimit(p, config.General.WindowSizeX, config.General.MarginOutsideScreen)) || (ParticleModification.ParticleOutsideYLimit(p, config.General.WindowSizeY, config.General.MarginOutsideScreen)) || p.Life <= 0 {
				s.KillParticule(particleNumber)
			}

			// Modifciation des parametre Couleur, Scale, Rotation, Opacity en fonction de mode choisie dans "config.json"
			var PositionMax int = config.General.WindowSizeX + config.General.WindowSizeY
			if config.General.ChangeColorAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
				p.ColorRed = ParticleModification.UpdateAccordingToPosition(p, config.General.MinColorRed, config.General.MaxColorRed, PositionMax)
				p.ColorBlue = ParticleModification.UpdateAccordingToPosition(p, config.General.MinColorBlue, config.General.MaxColorBlue, PositionMax)
				p.ColorGreen = ParticleModification.UpdateAccordingToPosition(p, config.General.MinColorGreen, config.General.MaxColorGreen, PositionMax)
			} else if config.General.ChangeColorAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
				p.ColorRed = ParticleModification.UpdateAccordingToLife(p, config.General.MinColorRed, config.General.MaxColorRed)
				p.ColorBlue = ParticleModification.UpdateAccordingToLife(p, config.General.MinColorBlue, config.General.MaxColorBlue)
				p.ColorGreen = ParticleModification.UpdateAccordingToLife(p, config.General.MinColorGreen, config.General.MaxColorGreen)
			}

			if config.General.ChangeScaleAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
				p.ScaleX = ParticleModification.UpdateAccordingToPosition(p, 0, config.General.ScaleX, PositionMax)
				p.ScaleY = ParticleModification.UpdateAccordingToPosition(p, 0, config.General.ScaleY, PositionMax)
			} else if config.General.ChangeScaleAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
				p.ScaleX = ParticleModification.UpdateAccordingToLife(p, 0, config.General.ScaleX)
				p.ScaleY = ParticleModification.UpdateAccordingToLife(p, 0, config.General.ScaleY)
			}

			if config.General.ChangeRotationAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
				p.Rotation = ParticleModification.UpdateAccordingToPosition(p, 0, config.General.Rotation, PositionMax)
			} else if config.General.ChangeRotationAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
				p.Rotation = ParticleModification.UpdateAccordingToLife(p, 0, config.General.Rotation)
			}

			if config.General.ChangeOpacityAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
				p.Opacity = ParticleModification.UpdateAccordingToPosition(p, 0, config.General.Opacity, PositionMax)
			} else if config.General.ChangeOpacityAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
				p.Opacity = ParticleModification.UpdateAccordingToLife(p, 0, config.General.Opacity)
			}
			//

			if config.General.Collision {
				if config.General.CollisionAmongParticle {
					CollisionAmongParticle(particleNumber, p)
				}
				if config.General.CollisionWithWall {
					CollisionWithWall(p, config.General.WindowSizeX, config.General.WindowSizeY)
				}
			}

		}
		particleNumber = particle
	}

	if config.General.Collision && config.General.CollisionAmongParticle {

		s.PlaceAccordingToPosition()

	}

	//ajout des particule
	if Part_particle >= 1 {
		for i := 0; float64(i) < Part_particle; i++ {
			Part_particle -= 1
			s.Add_Particule()
		}
	}
	//

	// Si SpanwCenter = true, on centre le generateur de particules
	if config.General.SpawnCenter {
		config.General.SpawnX = config.General.WindowSizeX / 2
		config.General.SpawnY = config.General.WindowSizeY / 2
	}

}
