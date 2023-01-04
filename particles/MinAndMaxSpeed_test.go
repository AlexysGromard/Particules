package particles

import "testing"

func TestSpeedMode0(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(0)
	if maxSpeed != 5 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrais être à 5 et 0 mais il sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestSpeedMode1(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(1)
	if maxSpeed != 1 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrais être à 1 et 0 mais il sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestSpeeMode2(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(2)
	if maxSpeed != 3 || minSpeed != 1 {
		t.Error("maxSpeed et minSpeed devrais être à 3 et 1 mais il sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestSpeeMode3(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(3)
	if maxSpeed != 5 || minSpeed != 3 {
		t.Error("maxSpeed et minSpeed devrais être à 5 et 3 mais il sont à ", maxSpeed, " et ", minSpeed)
	}
}
