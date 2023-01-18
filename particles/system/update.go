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
func (s *System) Update() {
	Part_particle += config.General.SpawnRate
	particleNumber := s.Content.Front()
	//actualisation des particules
	for particleNumber != nil {
		// on récupère la particule
		particule, ready := particleNumber.Value.(*particles.Particle)
		// on passe à la particule suivante
		nextParticle := particleNumber.Next()

		if ready {

			// mouvement des particule
			//la vitesse est modifier en fonction de l'accélération de la gravité
			particule.SpeedY = particule.SpeedY + config.General.Gravity
			//

			//la position est modifie ne fonction de la vitesse
			particule.PositionX = particule.PositionX + particule.SpeedX
			particule.PositionY = particule.PositionY + particule.SpeedY
			//
			//

			//quand le system de vie est active on vas retirer une uniter de vis à chaque update
			if config.General.HaveLife {
				particule.Life -= 1
			}
			//

			//tue les particules morte
			if (ParticleModification.ParticleOutsideXLimit(particule, config.General.WindowSizeX, config.General.MarginOutsideScreen)) || (ParticleModification.ParticleOutsideYLimit(particule, config.General.WindowSizeY, config.General.MarginOutsideScreen)) || particule.Life <= 0 {
				s.KillParticule(particleNumber)
			}

			// Modifciation des parametre Couleur, Scale, Rotation, Opacity en fonction de mode choisie dans "config.json"
			var PositionMax int = config.General.WindowSizeX + config.General.WindowSizeY
			if config.General.ChangeColorAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
				particule.ColorRed = ParticleModification.UpdateAccordingToPosition(particule, config.General.MinColorRed, config.General.MaxColorRed, PositionMax)
				particule.ColorBlue = ParticleModification.UpdateAccordingToPosition(particule, config.General.MinColorBlue, config.General.MaxColorBlue, PositionMax)
				particule.ColorGreen = ParticleModification.UpdateAccordingToPosition(particule, config.General.MinColorGreen, config.General.MaxColorGreen, PositionMax)
			} else if config.General.ChangeColorAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
				particule.ColorRed = ParticleModification.UpdateAccordingToLife(particule, config.General.MinColorRed, config.General.MaxColorRed)
				particule.ColorBlue = ParticleModification.UpdateAccordingToLife(particule, config.General.MinColorBlue, config.General.MaxColorBlue)
				particule.ColorGreen = ParticleModification.UpdateAccordingToLife(particule, config.General.MinColorGreen, config.General.MaxColorGreen)
			}

			if config.General.ChangeScaleAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
				particule.ScaleX = ParticleModification.UpdateAccordingToPosition(particule, 0, config.General.ScaleX, PositionMax)
				particule.ScaleY = ParticleModification.UpdateAccordingToPosition(particule, 0, config.General.ScaleY, PositionMax)
			} else if config.General.ChangeScaleAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
				particule.ScaleX = ParticleModification.UpdateAccordingToLife(particule, 0, config.General.ScaleX)
				particule.ScaleY = ParticleModification.UpdateAccordingToLife(particule, 0, config.General.ScaleY)
			}

			if config.General.ChangeRotationAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
				particule.Rotation = ParticleModification.UpdateAccordingToPosition(particule, 0, config.General.Rotation, PositionMax)
			} else if config.General.ChangeRotationAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
				particule.Rotation = ParticleModification.UpdateAccordingToLife(particule, 0, config.General.Rotation)
			}

			if config.General.ChangeOpacityAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
				particule.Opacity = ParticleModification.UpdateAccordingToPosition(particule, 0, config.General.Opacity, PositionMax)
			} else if config.General.ChangeOpacityAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
				particule.Opacity = ParticleModification.UpdateAccordingToLife(particule, 0, config.General.Opacity)
			}
			//

			if config.General.Collision {
				if config.General.CollisionAmongParticle {
					CollisionAmongParticle(particleNumber, particule)
				}
				if config.General.CollisionWithWall {
					CollisionWithWall(particule, config.General.WindowSizeX, config.General.WindowSizeY)
				}
			}

		}
		particleNumber = nextParticle
	}

	if config.General.Collision && config.General.CollisionAmongParticle {
		s.PlaceAccordingToPosition()
	}

	//ajout des particule
	if Part_particle >= 1 {
		nombreDeToureMax := Part_particle
		for nombreDeToure := 0; float64(nombreDeToure) < nombreDeToureMax; nombreDeToure++ {
			Part_particle -= 1
			s.AddParticule()
		}
	}
	//

	// Si SpanwCenter = true, on centre le generateur de particules
	if config.General.SpawnCenter {
		config.General.SpawnX = config.General.WindowSizeX / 2
		config.General.SpawnY = config.General.WindowSizeY / 2
	}

}
