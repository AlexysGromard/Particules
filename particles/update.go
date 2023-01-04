package particles

import (
	//"math/rand"
	//"math"
	"project-particles/config"
	//"time"
	//"fmt"
)

var Part_particle float64

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
			if !((p.PositionX < 0-config.General.MarginOutsideScreen || p.PositionX > float64(config.General.WindowSizeX)+config.General.MarginOutsideScreen) || (p.PositionY < 0-config.General.MarginOutsideScreen || p.PositionY > float64(config.General.WindowSizeY)+config.General.MarginOutsideScreen) || p.Life <= 0) {
				// mouvement des particule
				p.SpeedY = p.SpeedY + config.General.Gravity

				p.PositionX = p.PositionX + p.SpeedX
				p.PositionY = p.PositionY + p.SpeedY

				if config.General.HaveLife {
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

				var PositionMax int = config.General.WindowSizeX + config.General.WindowSizeY

				if config.General.ChangeColorAccordingTo == 1 {
					p.ColorRed = p.UpdateAccordingToPosition(0, config.General.ColorRed, PositionMax)
					p.ColorBlue = p.UpdateAccordingToPosition(0, config.General.ColorBlue, PositionMax)
					p.ColorGreen = p.UpdateAccordingToPosition(0, config.General.ColorGreen, PositionMax)
				} else if config.General.ChangeColorAccordingTo == 2 {
					p.ColorRed = p.UpdateAccordingToLife(0, config.General.ColorRed, config.General.Life)
					p.ColorBlue = p.UpdateAccordingToLife(0, config.General.ColorBlue, config.General.Life)
					p.ColorGreen = p.UpdateAccordingToLife(0, config.General.ColorGreen, config.General.Life)
				}

				if config.General.ChangeScaleAccordingTo == 1 {
					p.ScaleX = p.UpdateAccordingToPosition(0, config.General.ScaleX, PositionMax)
					p.ScaleX = p.UpdateAccordingToPosition(0, config.General.ScaleY, PositionMax)
				} else if config.General.ChangeScaleAccordingTo == 2 {
					p.ScaleX = p.UpdateAccordingToLife(0, config.General.ScaleX, config.General.Life)
					p.ScaleX = p.UpdateAccordingToLife(0, config.General.ScaleY, config.General.Life)
				}

				if config.General.ChangeRotationAccordingTo == 1 {
					p.Rotation = p.UpdateAccordingToPosition(0, config.General.Rotation, PositionMax)
				} else if config.General.ChangeRotationAccordingTo == 2 {
					p.Rotation = p.UpdateAccordingToLife(0, config.General.Rotation, config.General.Life)
				}

				if config.General.ChangeOpacityAccordingTo == 1 {
					p.Opacity = p.UpdateAccordingToPosition(0, config.General.Opacity, PositionMax)
				} else if config.General.ChangeOpacityAccordingTo == 2 {
					p.Opacity = p.UpdateAccordingToLife(0, config.General.Opacity, config.General.Life)
				}

				//enregistre les particule morte
				if (p.PositionX < 0-config.General.MarginOutsideScreen || p.PositionX > float64(config.General.WindowSizeX)+config.General.MarginOutsideScreen) || (p.PositionY < 0-config.General.MarginOutsideScreen || p.PositionY > float64(config.General.WindowSizeY)+config.General.MarginOutsideScreen) || p.Life <= 0 {
					DeadParticles.PushFront(p)
				}

				/*
					if config.General.Collision {
						var suivant *Particle = i.Next()

						if p.SpeedX < 0 && p.PositionX > suivant.Value.PositionX {
							s.Content.InsertBefore(suivant)
							s.Content.Remove(i)
							//p.PositionX, i.Next().Value(*Particle).PositionX = i.Next().Value(*Particle).PositionX, p.PositionX
						}

						var avant Particle = i.Pred()
						if p.SpeedX > 0 && p.PositionX > i.Prev().Value(*Particle).positionX {
							s.Content.InsertBefore(avant)
							s.Content.Remove(i)
							//p.PositionX, i.Prev().Value(*Particle).PositionX = i.Prev().Value(*Particle).PositionX, p.PositionX
						}

						for j := i.Next(); j != nil; j = j.Next() {
							q, ok2 := j.Value.(*Particle)
							if ok2 {
								if !((q.PositionX < 0-config.General.MarginOutsideScreen || q.PositionX > float64(config.General.WindowSizeX)+config.General.MarginOutsideScreen) || (q.PositionY < 0-config.General.MarginOutsideScreen || q.PositionY > float64(config.General.WindowSizeY)+config.General.MarginOutsideScreen) || q.Life <= 0) {
									if math.Abs(p.PositionX-q.PositionX) <= p.ScaleX*10 && math.Abs(p.PositionY-q.PositionY) <= p.ScaleY*10 && q != p {
										//q.ColorRed,q.ColorGreen,q.ColorBlue = 1,0,0
										if config.General.WhatCollisionDo == 1 {
											q.SpeedX, q.SpeedY, p.SpeedX, p.SpeedY = p.SpeedX, p.SpeedY, q.SpeedX, q.SpeedY
										}
									}
								}
							}
						}
					}*/
			}
		}
	}

	//ajout des particule
	if Part_particle >= 1 {

		for i := 0; float64(i) < Part_particle; i++ {
			Part_particle -= 1
			s.Add_Particule(DeadParticles)
		}
	}
	// Si SpanwCenter = true, on centre le generateur de particules
	if config.General.SpanwCenter && !config.General.Interaction {
		config.General.SpawnX = config.General.WindowSizeX / 2
		config.General.SpawnY = config.General.WindowSizeY / 2
	}
}
