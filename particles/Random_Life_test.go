package particles

import "testing"

func TestDansLesLimite(t *testing.T) {
	var vie int

	for i := 0; i < 100; i++ {
		vie = Random_Life()
		if vie < 0 || vie > 50 {
			t.Error("la vie générée aléatoirement devrait être dans l'intervalle [0; 50[ mais il elle est égale à ", vie)
			break
		}
	}
}
