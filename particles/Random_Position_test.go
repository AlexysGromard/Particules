package particles

import "testing"

func TestDansLInterval(t *testing.T) {
	var posX, posY float64

	for i := 0; i < 100; i++ {
		posX, posY = Random_Position(200, 100)

		if posX <= 0 || posX > 200 {
			t.Error("La position en X n'est pas dans l'interval [0;200[")
			break
		}
		if posY <= 0 || posY > 100 {
			t.Error("La position en Y n'est pas dans l'interval [0;100[")
			break
		}
	}
}
