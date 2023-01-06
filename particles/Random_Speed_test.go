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


func TestRandom_SpeedDirection(t *testing.T){
	var speedX,SpeedY float64

	//si une vitesse va de en haut à gauche alors BonneDirection[0] devrais être égale à
	//true de même si BonneDirection[1 ou 2 ou 3] si la particule vas en haut à droit ou en bas
	//à gauche ou en bas à droit 
	var BonneDirection [4]bool

	for i := 0; i<100 ; i++{
		speedX,SpeedY = Random_Speed(2)	

		if speedX > 0 && SpeedY <0{
			BonneDirection[0] = true
		}
		if speedX > 0 && SpeedY >0{
			BonneDirection[1] = true
		}
		if speedX < 0 && SpeedY <0{
			BonneDirection[2] = true
		}
		if speedX < 0 && SpeedY >0{
			BonneDirection[3] = true
		}
	}

	for _,valeur := range BonneDirection{
		if valeur == false{
			t.Error("la vitesse ne va dans tout les directions")
		}
	}
}