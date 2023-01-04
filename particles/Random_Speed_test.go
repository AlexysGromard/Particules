package particles

import (
	"testing"
	"math"
)

func TestRandom_Speed0(t *testing.T){
	if !Vérification(0){
		t.Error("les vitesse générée ne ne sont pas dans l'interval")
	}
}

func TestRandom_Speed1(t *testing.T){
	if !Vérification(1){
		t.Error("les vitesse générée ne ne sont pas dans l'interval")
	}
}

func TestRandom_Speed2(t *testing.T){
	if !Vérification(2){
		t.Error("les vitesse générée ne ne sont pas dans l'interval")
	}
}

func TestRandom_Speed3(t *testing.T){
	if !Vérification(3){
		t.Error("les vitesse générée ne ne sont pas dans l'interval")
	}
}


func Vérification(mode int) bool{
	var speedX,SpeedY,v_Réelle float64
	var min,max int

	max,min = MinAndMaxSpeed(mode)

	for i := 0; i<100 ; i++{
		speedX,SpeedY = Random_Speed(mode)	

		v_Réelle = math.Sqrt(carrée(speedX)+carrée(SpeedY))
	
		if v_Réelle <= float64(min) || v_Réelle > float64(max){
			return false
		}
	}

	return true
}

func carrée(n float64) float64{
	return n*n
}

