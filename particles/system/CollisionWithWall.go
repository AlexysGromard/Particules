package system

import (
	"project-particles/particles"
)

// La fonction CollisionWithWall établie les collisions avec les bords de l'écran
// Elle prend en paramètre une particule et les limites de l'écran
// Elle vérifie si la particule est assez proche des bords de l'écran pour qu'il y est une collision
// Si il y a une collision elle inverse la vitesse de la particule pour faire le rebond en fonction du bord où il y a la collision
func CollisionWithWall(ParticuleActuelle *particles.Particle, LimiteX, LimiteY int) {
	//collision avec le bord de droite ou de gauche
	if ParticuleActuelle.PositionX <= 0 || ParticuleActuelle.PositionX+ParticuleActuelle.ScaleX*10 >= float64(LimiteX) {
		//inverse la vitesse pour faire le rebon
		ParticuleActuelle.SpeedX = -ParticuleActuelle.SpeedX
		// Pour que les particules ne se retrouvent pas bloquées dans les bords de l'écran on les éloigne un peut du bord où il y a la collision
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
