package installationParticles

import "testing"

func TestDansLInterval1(t *testing.T) {
	var posX, posY float64
	var posXMax, posYMax int = 100, 75

	for i := 0; i < 100; i++ {
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

func TestDansLInterval2(t *testing.T) {
	var posX, posY float64
	var posXMax, posYMax int = 100, 75

	for i := 0; i < 100; i++ {
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
