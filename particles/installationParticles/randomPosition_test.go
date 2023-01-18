package installationParticles

import "testing"

// Le test TestDansLInterval1 vérifie que la fonction Random_Position retourne bien des positions
// dans l'intervalle [0; 100[ en X et [0; 75[ en Y
// Si ce n'est pas le cas, le test échoue
func TestDansLInterval1(t *testing.T) {
	var posX, posY float64
	var posXMax, posYMax int = 100, 75

	for NuméroToure := 0; NuméroToure < 100; NuméroToure++ {
		posX, posY = Random_Position(posXMax, posYMax)

		if posX <= 0 || posX > float64(posXMax) {
			t.Error("La position en X n'est pas dans l'intervalle [0;200[")
			break
		}
		if posY <= 0 || posY > float64(posYMax) {
			t.Error("La position en Y n'est pas dans l'intervalle [0;100[")
			break
		}
	}
}

// Le test TestDansLInterval2 vérifie que la fonction Random_Position retourne bien des positions
// dans l'intervalle [0; 200[ en X et [0; 100[ en Y
// Si ce n'est pas le cas, le test échoue
func TestDansLInterval2(t *testing.T) {
	var posX, posY float64
	var posXMax, posYMax int = 100, 75

	for NuméroToure := 0; NuméroToure < 100; NuméroToure++ {
		posX, posY = Random_Position(posXMax, posYMax)

		if posX <= 0 || posX > float64(posXMax) {
			t.Error("La position en X n'est pas dans l'intervalle [0;200[")
			break
		}
		if posY <= 0 || posY > float64(posYMax) {
			t.Error("La position en Y n'est pas dans l'intervalle [0;100[")
			break
		}
	}
}
