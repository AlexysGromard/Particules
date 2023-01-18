package system

import (
	"project-particles/particles"
)

func CollisionWithWall(ParticuleActuelle *particles.Particle, LimiteX, LimiteY int) {
	//collision avec le bord de droite ou de gauche
	if ParticuleActuelle.PositionX <= 0 || ParticuleActuelle.PositionX+ParticuleActuelle.ScaleX*10 >= float64(LimiteX) {
		//inverse la vitesse pour faire le rebon
		ParticuleActuelle.SpeedX = -ParticuleActuelle.SpeedX

		// pour que les particules ne se retrouvent pas bloquées dans les bords de l'écran on les éloigne un peut du bord où il y a la collision
		if ParticuleActuelle.PositionX <= 0 {
			ParticuleActuelle.PositionX = ParticuleActuelle.PositionX + ParticuleActuelle.ScaleX*0.1
		} else {
			ParticuleActuelle.PositionX = ParticuleActuelle.PositionX - ParticuleActuelle.ScaleX*0.1
		}
	}

	//collision avec le bord de haut ou de bas
	if ParticuleActuelle.PositionY <= 0 || ParticuleActuelle.PositionY+ParticuleActuelle.ScaleY*10 >= float64(LimiteY) {
		//inverse la vitesse pour faire le rebon
		ParticuleActuelle.SpeedY = -ParticuleActuelle.SpeedY
		// pour que les particules ne se retrouvent pas bloquées dans les bords de l'écran on les éloigne un peut du bord où il y a la collision
		if ParticuleActuelle.PositionY <= 0 {
			ParticuleActuelle.PositionY = ParticuleActuelle.PositionY + ParticuleActuelle.ScaleY*0.1
		} else {
			ParticuleActuelle.PositionY = ParticuleActuelle.PositionY - ParticuleActuelle.ScaleY*0.1
		}
	}
}
