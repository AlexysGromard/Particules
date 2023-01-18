package installationParticles

import "testing"

func TestDansLesLimite1(t *testing.T) {
	var vie int
	var vieMax uint = 50

	for i := 0; i < 100; i++ {
		vie = Random_Life(vieMax)
		if vie < 0 || vie > int(vieMax) {
			t.Error("la vie générée aléatoirement devrait être dans l'intervalle [0; 50[ mais il elle est égale à ", vie)
			break
		}
	}
}

func TestDansLesLimite2(t *testing.T) {
	var vie int
	var vieMax uint = 100

	for i := 0; i < 100; i++ {
		vie = Random_Life(vieMax)
		if vie < 0 || vie > int(vieMax) {
			t.Error("la vie générée aléatoirement devrait être dans l'intervalle [0; 50[ mais il elle est égale à ", vie)
			break
		}
	}
}
