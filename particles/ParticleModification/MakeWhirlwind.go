package ParticleModification

import (
	"container/list"
	"math"
	"project-particles/particles"
)

// La fonction MakeWhirlwind crée un tourbillon de particules
// Elle est appelée grâce à la touche T
func MakeWhirlwind(list *list.List, SpeedType int, centreX, centreY float64) {
	for i := list.Front(); i != nil; i = i.Next() {
		p, ok := i.Value.(*particles.Particle)
		if ok {
			//Calcule de la distance par raport au centre
			distance := math.Sqrt((p.PositionX-centreX)*(p.PositionX-centreX) + (p.PositionY-centreY)*(p.PositionY-centreY))

			//pour éviter de stoper les particules au moment de leur aparition au centre On applique l'effect,
			//uniquement quand elle ne sont plus au centre donc que la distance est égale à 0
			if distance != 0 {
				// Calcul de l'angle actuel de la particule par rapport au générateur
				angle := math.Atan2(p.PositionY-float64(centreY), p.PositionX-float64(centreX))
				// Dela définit la vitesse de rotation
				delta := 0.7
				angle += delta

				// On recalcule les composantes de vitesse en utilisant les nouvelles valeurs d'angle
				p.SpeedX = distance * math.Cos(angle)
				p.SpeedY = distance * math.Sin(angle)

				// speedFactor ralenti la particule en fonction du SpeedType
				speedFactor := 0.04 * float64(SpeedType)
				p.SpeedX = p.SpeedX * speedFactor
				p.SpeedY = p.SpeedY * speedFactor
			}

		}

	}

}
