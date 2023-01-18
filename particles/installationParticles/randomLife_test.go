package installationParticles

import "testing"

// Le test TestDansLesLimite1 vérifie que la fonction Random_Life retourne bien une valeur dans l'intervalle [0; 50[
// Si ce n'est pas le cas, le test échoue
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

// Le test TestDansLesLimite2 vérifie que la fonction Random_Life retourne bien une valeur dans l'intervalle [0; 100[
// Si ce n'est pas le cas, le test échoue
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
