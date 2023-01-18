package installationParticles

import "testing"

// Le test TestSpeedModemoin1 vérifie que la fonction minAndMaxSpeed retourne bien 0 et 5
// lorsque le mode de vitesse est -1
// Si ce n'est pas le cas, le test échoue
func TestSpeedModemoin1(t *testing.T) {
	maxSpeed, minSpeed := minAndMaxSpeed(-1)
	if maxSpeed != 5 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrait être à 5 et 0 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

// Le test TestSpeedMode1 vérifie que la fonction minAndMaxSpeed retourne bien 1 et 0
// lorsque le mode de vitesse est 1
// Si ce n'est pas le cas, le test échoue
func TestSpeedMode1(t *testing.T) {
	maxSpeed, minSpeed := minAndMaxSpeed(1)
	if maxSpeed != 1 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrait être à 1 et 0 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

// Le test TestSpeeMode2 vérifie que la fonction minAndMaxSpeed retourne bien 3 et 1
// lorsque le mode de vitesse est 2
// Si ce n'est pas le cas, le test échoue
func TestSpeeMode2(t *testing.T) {
	maxSpeed, minSpeed := minAndMaxSpeed(2)
	if maxSpeed != 3 || minSpeed != 1 {
		t.Error("maxSpeed et minSpeed devrait être à 3 et 1 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

// Le test TestSpeeMode3 vérifie que la fonction minAndMaxSpeed retourne bien 5 et 3
// lorsque le mode de vitesse est 3
// Si ce n'est pas le cas, le test échoue
func TestSpeeMode3(t *testing.T) {
	maxSpeed, minSpeed := minAndMaxSpeed(3)
	if maxSpeed != 5 || minSpeed != 3 {
		t.Error("maxSpeed et minSpeed devrait être à 5 et 3 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

// Le test TestSpeeModeOther1 vérifie que la fonction minAndMaxSpeed retourne bien -1 et -1
// lorsque le mode de vitesse est 10
// Si ce n'est pas le cas, le test échoue
func TestSpeeModeOther1(t *testing.T) {
	maxSpeed, minSpeed := minAndMaxSpeed(10)
	if maxSpeed != -1 || minSpeed != -1 {
		t.Error("maxSpeed et minSpeed devrait être à 5 et 3 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

// Le test TestSpeeModeOther2 vérifie que la fonction minAndMaxSpeed retourne bien -1 et -1
// lorsque le mode de vitesse est 0
// Si ce n'est pas le cas, le test échoue
func TestSpeeModeOther2(t *testing.T) {
	maxSpeed, minSpeed := minAndMaxSpeed(0)
	if maxSpeed != -1 || minSpeed != -1 {
		t.Error("maxSpeed et minSpeed devrait être à -1 et -1 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}
