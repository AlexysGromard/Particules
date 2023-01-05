package particles

import (
	"testing"
)

func Test1AccordingToLife(t *testing.T) {
	var particule Particle
	var minCouleur, maxCouleur float64 = 0, 1

	particule = Basique_Particule()
	couleur_init := particule.UpdateAccordingToLife(minCouleur, maxCouleur)

	particule.Life = 10
	couleur_fin := particule.UpdateAccordingToLife(minCouleur, maxCouleur)

	if couleur_init == couleur_fin {
		t.Error("la couleur n'a pas changer alors que la position à changer")
	}
}

func Test2AccordingToLife(t *testing.T) {
	var particule Particle
	var minCouleur, maxCouleur float64 = 0, 1

	particule = Basique_Particule()
	particule.Life = 25
	couleur := particule.UpdateAccordingToLife(minCouleur, maxCouleur)

	if couleur != 0.5 {
		t.Error("la couleur n'a pas changer alors que la position à changer")
	}
}

func Basique_Particule() Particle {
	return Particle{
		PositionX: 0,
		PositionY: 0,
		Rotation:  0,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:  1,
		SpeedX:   1,
		SpeedY:   1,
		LifeInit: 50,
		Life:     50,
	}
}
