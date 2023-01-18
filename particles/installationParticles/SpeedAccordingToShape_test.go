package installationParticles

import (
	"testing"
	//"project-particles/particles"
	"project-particles/particles/Test"
)

// Le test TestSpeedAccordingToShape1 vérifie que la fonction SpeedAccordingToShape renvoie une vitesse positive si la particule est dans le bon sens
// Si ce n'est pas le cas, le test échoue
func TestSpeedAccordingToShape1(t *testing.T) {
	centreX, centreY := float64(100), float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 110, 110

	speedX, speedY := SpeedAccordingToShape(1, particle.PositionX, particle.PositionY, centreX, centreY)

	if speedX < 0 || speedY < 0 {
		t.Error("la particule ne va pas dans la bonne direction")
	}
}

// Le test TestSpeedAccordingToShape2 vérifie que la fonction SpeedAccordingToShape renvoie une vitesse positive si la particule est dans le bon sens
// Si ce n'est pas le cas, le test échoue
func TestSpeedAccordingToShape2(t *testing.T) {
	centreX, centreY := float64(100), float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 110, 90

	speedX, speedY := SpeedAccordingToShape(1, particle.PositionX, particle.PositionY, centreX, centreY)

	if speedX < 0 || speedY > 0 {
		t.Error("la particule ne va pas dans la bonne direction")
	}
}

// Le test TestSpeedAccordingToShape3 vérifie que la fonction SpeedAccordingToShape renvoie une vitesse positive si la particule est dans le bon sens
// Si ce n'est pas le cas, le test échoue
func TestSpeedAccordingToShape3(t *testing.T) {
	centreX, centreY := float64(100), float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 90, 90

	speedX, speedY := SpeedAccordingToShape(1, particle.PositionX, particle.PositionY, centreX, centreY)

	if speedX > 0 || speedY > 0 {
		t.Error("la particule ne va pas dans la bonne direction")
	}
}

// Le test TestSpeedAccordingToShape4 vérifie que la fonction SpeedAccordingToShape renvoie une vitesse positive si la particule est dans le bon sens
// Si ce n'est pas le cas, le test échoue
func TestSpeedAccordingToShape4(t *testing.T) {
	centreX, centreY := float64(100), float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 90, 110

	speedX, speedY := SpeedAccordingToShape(1, particle.PositionX, particle.PositionY, centreX, centreY)

	if speedX > 0 || speedY < 0 {
		t.Error("la particule ne va pas dans la bonne direction")
	}
}

// Le test TestSpeedAccordingToShapeSpeed0 vérifie que la fonction SpeedAccordingToShape renvoie une vitesse nulle si la particule est dans le bon sens
func TestSpeedAccordingToShapeSpeed0(t *testing.T) {
	centreX, centreY := float64(100), float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 110, 110

	speedX, speedY := SpeedAccordingToShape(0, particle.PositionX, particle.PositionY, centreX, centreY)

	if speedX != 0 || speedY != 0 {
		t.Error("la particule ne va pas ces vitesse à 0")
	}
}
