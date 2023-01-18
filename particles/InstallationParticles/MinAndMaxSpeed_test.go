package InstallationParticles

import "testing"

func TestSpeedModemoin1(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(-1)
	if maxSpeed != 5 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrait être à 5 et 0 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestSpeedMode1(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(1)
	if maxSpeed != 1 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrait être à 1 et 0 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestSpeeMode2(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(2)
	if maxSpeed != 3 || minSpeed != 1 {
		t.Error("maxSpeed et minSpeed devrait être à 3 et 1 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestSpeeMode3(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(3)
	if maxSpeed != 5 || minSpeed != 3 {
		t.Error("maxSpeed et minSpeed devrait être à 5 et 3 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestSpeeModeOther1(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(10)
	if maxSpeed != -1 || minSpeed != -1 {
		t.Error("maxSpeed et minSpeed devrait être à 5 et 3 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestSpeeModeOther2(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(0)
	if maxSpeed != -1 || minSpeed != -1 {
		t.Error("maxSpeed et minSpeed devrait être à 5 et 3 mais ils sont à ", maxSpeed, " et ", minSpeed)
	}
}
