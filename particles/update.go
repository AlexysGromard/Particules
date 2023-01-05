package particles

import (
	"project-particles/config"
)

var Part_particle float64 // cette variable corespon au partie de particule

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
				//la vitesse est modifier en fonction de l'accélération de la gravité
				p.SpeedY = p.SpeedY + config.General.Gravity

				//la position est modifie ne fonction de la vitesse
				p.PositionX = p.PositionX + p.SpeedX
				p.PositionY = p.PositionY + p.SpeedY

				//quand le system de vie est active on vas retirer une uniter de vis à chaque update
				if config.General.HaveLife {
					p.Life -= 1
				}

				var PositionMax int = config.General.WindowSizeX + config.General.WindowSizeY
				if config.General.ChangeColorAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
					p.ColorRed = p.UpdateAccordingToPosition(0, config.General.ColorRed, PositionMax)
					p.ColorBlue = p.UpdateAccordingToPosition(0, config.General.ColorBlue, PositionMax)
					p.ColorGreen = p.UpdateAccordingToPosition(0, config.General.ColorGreen, PositionMax)
				} else if config.General.ChangeColorAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
					p.ColorRed = p.UpdateAccordingToLife(0, config.General.ColorRed)
					p.ColorBlue = p.UpdateAccordingToLife(0, config.General.ColorBlue)
					p.ColorGreen = p.UpdateAccordingToLife(0, config.General.ColorGreen)
				}

				if config.General.ChangeScaleAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
					p.ScaleX = p.UpdateAccordingToPosition(0, config.General.ScaleX, PositionMax)
					p.ScaleX = p.UpdateAccordingToPosition(0, config.General.ScaleY, PositionMax)
				} else if config.General.ChangeScaleAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
					p.ScaleX = p.UpdateAccordingToLife(0, config.General.ScaleX)
					p.ScaleX = p.UpdateAccordingToLife(0, config.General.ScaleY)
				}

				if config.General.ChangeRotationAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
					p.Rotation = p.UpdateAccordingToPosition(0, config.General.Rotation, PositionMax)
				} else if config.General.ChangeRotationAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
					p.Rotation = p.UpdateAccordingToLife(0, config.General.Rotation)
				}

				if config.General.ChangeOpacityAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
					p.Opacity = p.UpdateAccordingToPosition(0, config.General.Opacity, PositionMax)
				} else if config.General.ChangeOpacityAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
					p.Opacity = p.UpdateAccordingToLife(0, config.General.Opacity)
				}

				//enregistre les particule morte
				if (p.PositionX < 0-config.General.MarginOutsideScreen || p.PositionX > float64(config.General.WindowSizeX)+config.General.MarginOutsideScreen) || (p.PositionY < 0-config.General.MarginOutsideScreen || p.PositionY > float64(config.General.WindowSizeY)+config.General.MarginOutsideScreen) || p.Life <= 0 {
					DeadParticles.PushFront(p)
				}
				/*
					if config.General.Collision {
						//suivant := i.Next()
						//particule_suivante, _ := suivant.Value.(*Particle)
						var ok4, ok3 bool = true, true
						var V1, V2 *Particle

						fmt.Println(s.Content)
						fmt.Println(i)

						V1, ok4 = i.Next().Value.(*Particle)
						fmt.Println("test1")
						if Premier_Tour == true {
							V2, ok3 = i.Prev().Value.(*Particle)
						}

						if ok4 && ok3 && Premier_Tour {
							if p.SpeedX < 0 && p.PositionX > V1.PositionX { // i.Next().Value.(*Particle).PositionX
								s.Content.InsertBefore(i, i.Next())
								s.Content.Remove(i)
							} else if p.SpeedX > 0 && p.PositionX > V2.PositionX { // i.Prev().Value.(*Particle).PositionX
								s.Content.InsertBefore(i, i.Prev())
								s.Content.Remove(i)
							}
							fmt.Println("test2")
							for j := i.Next(); j != nil; j = j.Next() {
								fmt.Println("test3")
								q, ok2 := j.Value.(*Particle)
								if ok2 {
									if !(q.PositionX < 0-config.General.MarginOutsideScreen || q.PositionX > float64(config.General.WindowSizeX)+config.General.MarginOutsideScreen) {
										if !((q.PositionY < 0-config.General.MarginOutsideScreen || q.PositionY > float64(config.General.WindowSizeY)+config.General.MarginOutsideScreen) || q.Life <= 0) {
											if math.Abs(p.PositionX-q.PositionX) <= p.ScaleX*10 && math.Abs(p.PositionY-q.PositionY) <= p.ScaleY*10 && q != p {
												//q.ColorRed,q.ColorGreen,q.ColorBlue = 1,0,0
												if config.General.WhatCollisionDo == 1 {
													q.SpeedX, q.SpeedY, p.SpeedX, p.SpeedY = p.SpeedX, p.SpeedY, q.SpeedX, q.SpeedY
												}
											}
										}
									} else {
										break
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
