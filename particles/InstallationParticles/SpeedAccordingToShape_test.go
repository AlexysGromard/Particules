package InstallationParticles

import (
	"testing"
	//"project-particles/particles"
	"project-particles/particles/Test"

)

func TestSpeedAccordingToShape1(t *testing.T){	
	centreX,centreY := float64(100),float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 110,110

	speedX,speedY := SpeedAccordingToShape(1, particle.PositionX, particle.PositionY, centreX,centreY )

	if speedX < 0 || speedY < 0{
		t.Error("la particule ne va pas dans la bonne direction")
	}
}

func TestSpeedAccordingToShape2(t *testing.T){	
	centreX,centreY := float64(100),float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 110,90

	speedX,speedY := SpeedAccordingToShape(1, particle.PositionX, particle.PositionY, centreX,centreY )

	if speedX < 0 || speedY > 0{
		t.Error("la particule ne va pas dans la bonne direction")
	}
}

func TestSpeedAccordingToShape3(t *testing.T){	
	centreX,centreY := float64(100),float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 90,90

	speedX,speedY := SpeedAccordingToShape(1, particle.PositionX, particle.PositionY, centreX,centreY )

	if speedX > 0 || speedY > 0{
		t.Error("la particule ne va pas dans la bonne direction")
	}
}

func TestSpeedAccordingToShape4(t *testing.T){	
	centreX,centreY := float64(100),float64(100)

	particle := Test.Basique_Particule()

	particle.PositionX, particle.PositionY = 90,110

	speedX,speedY := SpeedAccordingToShape(1, particle.PositionX, particle.PositionY, centreX,centreY )

	if speedX > 0 || speedY < 0{
		t.Error("la particule ne va pas dans la bonne direction")
	}
}