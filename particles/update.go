package particles

import (
	"math"
	"project-particles/config"
)

var Part_particle float64       // cette variable correspond au partie de particule
var WhirlwindState bool = false // cette variable permet de savoir si le tourbillon est actif ou non

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	Part_particle += config.General.SpawnRate

	i := s.Content.Front()

	//actualisation des particules
	for i != nil {
		p, ok := i.Value.(*Particle)

		suivant := i.Next()

		if ok {
			if !((p.PositionX < 0-config.General.MarginOutsideScreen || p.PositionX > float64(config.General.WindowSizeX)+config.General.MarginOutsideScreen) || (p.PositionY < 0-config.General.MarginOutsideScreen || p.PositionY > float64(config.General.WindowSizeY)+config.General.MarginOutsideScreen) || p.Life <= 0) {
				// mouvement des particule
				//la vitesse est modifier en fonction de l'accélération de la gravité
				p.SpeedY = p.SpeedY + config.General.Gravity
				//

				//la position est modifie ne fonction de la vitesse
				p.PositionX = p.PositionX + p.SpeedX
				p.PositionY = p.PositionY + p.SpeedY
				//

				//quand le system de vie est active on vas retirer une uniter de vis à chaque update
				if config.General.HaveLife {
					p.Life -= 1
				}
				//

				// Modifciation des parametre Couleur, Scale, Rotation, Opacity en fonction de mode choisie dans "config.json"
				var PositionMax int = config.General.WindowSizeX + config.General.WindowSizeY
				if config.General.ChangeColorAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
					p.ColorRed = p.UpdateAccordingToPosition(config.General.MinColorRed, config.General.MaxColorRed, PositionMax)
					p.ColorBlue = p.UpdateAccordingToPosition(config.General.MinColorBlue, config.General.MaxColorBlue, PositionMax)
					p.ColorGreen = p.UpdateAccordingToPosition(config.General.MinColorGreen, config.General.MaxColorGreen, PositionMax)
				} else if config.General.ChangeColorAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
					p.ColorRed = p.UpdateAccordingToLife(config.General.MinColorRed, config.General.MaxColorRed)
					p.ColorBlue = p.UpdateAccordingToLife(config.General.MinColorBlue, config.General.MaxColorBlue)
					p.ColorGreen = p.UpdateAccordingToLife(config.General.MinColorGreen, config.General.MaxColorGreen)
				}

				if config.General.ChangeScaleAccordingTo == 1 { //modifie les parametre en fonction de la position de la particule
					p.ScaleX = p.UpdateAccordingToPosition(0, config.General.ScaleX, PositionMax)
					p.ScaleY = p.UpdateAccordingToPosition(0, config.General.ScaleY, PositionMax)
				} else if config.General.ChangeScaleAccordingTo == 2 { //modifie les parametre en fonction de la position de la vie
					p.ScaleX = p.UpdateAccordingToLife(0, config.General.ScaleX)
					p.ScaleY = p.UpdateAccordingToLife(0, config.General.ScaleY)
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
				//

				//enregistre les particule morte
				if (p.PositionX < 0-config.General.MarginOutsideScreen || p.PositionX > float64(config.General.WindowSizeX)+config.General.MarginOutsideScreen) || (p.PositionY < 0-config.General.MarginOutsideScreen || p.PositionY > float64(config.General.WindowSizeY)+config.General.MarginOutsideScreen) || p.Life <= 0 {
					s.KillParticule(i)
				}

				if config.General.Collision {
					for j := i.Next(); j != nil; j = j.Next() {
						q, ok2 := j.Value.(*Particle)
						if ok2 {
							if !(q.PositionX < 0-config.General.MarginOutsideScreen || q.PositionX > float64(config.General.WindowSizeX)+config.General.MarginOutsideScreen) {
								if !((q.PositionY < 0-config.General.MarginOutsideScreen || q.PositionY > float64(config.General.WindowSizeY)+config.General.MarginOutsideScreen) || q.Life <= 0) {
									if math.Abs(p.PositionX-q.PositionX) <= p.ScaleX*10 && math.Abs(p.PositionY-q.PositionY) <= p.ScaleY*10 && q != p {
										//q.ColorRed,q.ColorGreen,q.ColorBlue = 1,0,0
										if config.General.WhatCollisionDo == 1 {
											q.SpeedX, q.SpeedY, p.SpeedX, p.SpeedY = p.SpeedX, p.SpeedY, q.SpeedX, q.SpeedY
										}

										angle1 := math.Atan2(p.PositionY-q.PositionY, p.PositionX-q.PositionX)

										p.PositionX = p.PositionX + math.Cos(angle1)*p.ScaleX*0.1
										p.PositionY = p.PositionY + math.Sin(angle1)*p.ScaleY*0.1
										q.PositionX = q.PositionX + math.Cos(-angle1)*q.ScaleX*0.1
										q.PositionY = q.PositionY + math.Sin(-angle1)*q.ScaleY*0.1
									}
								}
							} else {
								break
							}
						}
					}
					if p.PositionX <= 0 || p.PositionX+p.ScaleX*10 >= float64(config.General.WindowSizeX) {
						p.SpeedX = -p.SpeedX
						if p.PositionX <= 0 {
							p.PositionX = p.PositionX + p.ScaleX*0.1
						} else {
							p.PositionX = p.PositionX - p.ScaleX*0.1
						}
					}
					if p.PositionY <= 0 || p.PositionY+p.ScaleY*10 >= float64(config.General.WindowSizeY) {
						p.SpeedY = -p.SpeedY
						if p.PositionY <= 0 {
							p.PositionY = p.PositionY + p.ScaleY*0.1
						} else {
							p.PositionY = p.PositionY - p.ScaleY*0.1
						}
					}
				}
				// Tourbillon
				if WhirlwindState {
					p.SpeedX, p.SpeedY = MakeWhirlwind(p)
				}
			}
		}
		i = suivant
	}

	//trier des particules
	for i := s.Content.Front(); i != nil; i = i.Next() {
		s.PlaceAccordingToPosition(i)
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
	if config.General.SpanwCenter && !config.General.Interaction {
		config.General.SpawnX = config.General.WindowSizeX / 2
		config.General.SpawnY = config.General.WindowSizeY / 2
	}

}
