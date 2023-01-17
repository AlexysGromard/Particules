package ParticleModification

import (
	"math"
	"project-particles/config"
	"project-particles/particles"
	"container/list"

)

func MakeWhirlwind(list *list.List){//p *particles.Particle) {

	for i := list.Front(); i != nil ; i = i.Next(){
		p,ok := i.Value.(*particles.Particle)
		if ok{
			distance := math.Sqrt(math.Pow(p.PositionX-float64(config.General.SpawnX), 2) + math.Pow(p.PositionY-float64(config.General.SpawnY), 2))
			// Calcul de l'angle actuel de la particule par rapport au générateur
			angle := math.Atan2(p.PositionY-float64(config.General.SpawnY), p.PositionX-float64(config.General.SpawnX))
			// Dela définit la vitesse de rotation
			delta := 0.7
			angle += delta
		
			// On recalcule les composantes de vitesse en utilisant les nouvelles valeurs d'angle
			p.SpeedX = distance * math.Cos(angle)
			p.SpeedY = distance * math.Sin(angle)
			// speedFactor ralenti la particule
			speedFactor := 0.04 * float64(config.General.SpeedType)
			p.SpeedX = p.SpeedX * speedFactor
			p.SpeedY = p.SpeedY * speedFactor
		}

	}
	

}
