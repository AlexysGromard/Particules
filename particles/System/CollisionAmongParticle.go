package System

import (
	"container/list"
	"math"
	"project-particles/particles"
)

func CollisionAmongParticle(ElementActuel *list.Element, ParticuleActuelle *particles.Particle) {

	for ElementProche := ElementActuel.Next(); ElementProche != nil; ElementProche = ElementProche.Next() {
		ParticuleProche, ok2 := ElementProche.Value.(*particles.Particle)
		if ok2 {
			if math.Abs(ParticuleActuelle.PositionX-ParticuleProche.PositionX) <= ParticuleActuelle.ScaleX*10 {
				if math.Abs(ParticuleActuelle.PositionY-ParticuleProche.PositionY) <= ParticuleActuelle.ScaleY*10 && ParticuleProche != ParticuleActuelle {
					// quand il y a une collision on échange les vitesses
					ParticuleProche.SpeedX, ParticuleProche.SpeedY, ParticuleActuelle.SpeedX, ParticuleActuelle.SpeedY = ParticuleActuelle.SpeedX, ParticuleActuelle.SpeedY, ParticuleProche.SpeedX, ParticuleProche.SpeedY
					// pour éviter que des particules ne se retrouve collée on les éloigne un peu l'une de l'autre
					angle1 := math.Atan2(ParticuleActuelle.PositionY-ParticuleProche.PositionY, ParticuleActuelle.PositionX-ParticuleProche.PositionX)

					ParticuleActuelle.PositionX = ParticuleActuelle.PositionX + math.Cos(angle1)*ParticuleActuelle.ScaleX*0.1
					ParticuleActuelle.PositionY = ParticuleActuelle.PositionY + math.Sin(angle1)*ParticuleActuelle.ScaleY*0.1
					ParticuleProche.PositionX = ParticuleProche.PositionX + math.Cos(-angle1)*ParticuleProche.ScaleX*0.1
					ParticuleProche.PositionY = ParticuleProche.PositionY + math.Sin(-angle1)*ParticuleProche.ScaleY*0.1
				}
			} else {
				break
			}
		}
	}
}