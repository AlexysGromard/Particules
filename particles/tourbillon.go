package particles

import (
	"math"
	"project-particles/config"
)

func MakeWhirlwind(particle *Particle) (float64, float64) {
	var speedX, speedY float64
	distance := math.Sqrt(math.Pow(particle.PositionX-float64(config.General.SpawnX), 2) + math.Pow(particle.PositionY-float64(config.General.SpawnY), 2))
	// Calcul de l'angle actuel de la particule par rapport au générateur
	angle := math.Atan2(particle.PositionY-float64(config.General.SpawnY), particle.PositionX-float64(config.General.SpawnX))
	// Dela définit la vitesse de rotation
	delta := 0.7
	angle += delta

	// On recalcule les composantes de vitesse en utilisant les nouvelles valeurs d'angle
	speedX = distance * math.Cos(angle)
	speedY = distance * math.Sin(angle)
	// speedFactor ralenti la particule
	speedFactor := 0.04 * float64(config.General.SpeedType)
	speedX = speedX * speedFactor
	speedY = speedY * speedFactor
	return speedX, speedY
}
