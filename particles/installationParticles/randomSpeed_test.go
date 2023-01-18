package installationParticles

import (
	"math"
	"testing"
)

// Le test TestRandom_Speedmoin1 vérifie que la fonction randomSpeed retourne bien des vitesses comprises entre 0 et 5
// lorsque le mode de vitesse est -1
// Si ce n'est pas le cas, le test échoue
func TestRandom_Speedmoin1(t *testing.T) {
	if !Vérification(-1) {
		t.Error("Les vitesses générées ne sont pas dans l'intervalle")
	}
}

// Le test TestRandom_Speed0 vérifie que la fonction randomSpeed retourne bien uen vitesse nulle
// lorsque le mode de vitesse est 0
// Si ce n'est pas le cas, le test échoue
func TestRandom_Speed0(t *testing.T) {
	if !Vérification(0) {
		t.Error("Les vitesses générées ne sont pas dans l'intervalle")
	}
}

// Le test TestRandom_Speed1 vérifie que la fonction randomSpeed retourne bien des vitesses comprises entre 0 et 1
// lorsque le mode de vitesse est 1
// Si ce n'est pas le cas, le test échoue
func TestRandom_Speed1(t *testing.T) {
	if !Vérification(1) {
		t.Error("Les vitesses générées ne sont pas dans l'intervalle")
	}
}

// Le test TestRandom_Speed2 vérifie que la fonction randomSpeed retourne bien des vitesses comprises entre 1 et 3
// lorsque le mode de vitesse est 2
// Si ce n'est pas le cas, le test échoue
func TestRandom_Speed2(t *testing.T) {
	if !Vérification(2) {
		t.Error("Les vitesses générées ne sont pas dans l'intervalle")
	}
}

// Le test TestRandom_Speed3 vérifie que la fonction randomSpeed retourne bien des vitesses comprises entre 3 et 5
// lorsque le mode de vitesse est 3
// Si ce n'est pas le cas, le test échoue
func TestRandom_Speed3(t *testing.T) {
	if !Vérification(3) {
		t.Error("Les vitesses générées ne sont pas dans l'intervalle")
	}
}

func Vérification(mode int) bool {
	var speedX, SpeedY, v_Réelle float64
	var min, max int

	max, min = minAndMaxSpeed(mode)

	for i := 0; i < 100; i++ {
		speedX, SpeedY = Random_Speed(mode)

		v_Réelle = math.Sqrt(carrée(speedX) + carrée(SpeedY)) //la fonction Sqrt du math accepte considère que Sqrt(0) == 0

		if mode != 0 && (v_Réelle <= float64(min) || v_Réelle >= float64(max)) {
			return false
		}
	}

	return true
}

func carrée(n float64) float64 {
	return n * n
}

func TestRandom_SpeedDirection(t *testing.T) {
	var speedX, SpeedY float64

	//si une vitesse va de en haut à gauche alors BonneDirection[0] devrait être égale à
	//true de même si BonneDirection[1 ou 2 ou 3] si la particule va en haut à droit ou en bas
	//à gauche ou en bas à droit
	var BonneDirection [4]bool

	for i := 0; i < 100; i++ {
		speedX, SpeedY = Random_Speed(2)

		if speedX > 0 && SpeedY < 0 {
			BonneDirection[0] = true
		}
		if speedX > 0 && SpeedY > 0 {
			BonneDirection[1] = true
		}
		if speedX < 0 && SpeedY < 0 {
			BonneDirection[2] = true
		}
		if speedX < 0 && SpeedY > 0 {
			BonneDirection[3] = true
		}
	}

	for _, valeur := range BonneDirection {
		if valeur == false {
			t.Error("la vitesse ne va pas dans toutes les directions")
		}
	}
}
