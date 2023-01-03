package particles

import "testing"

func TestMode0(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(0)
	if maxSpeed != 5 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrais être à 5 et 0 mais il sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestMode1(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(0)
	if maxSpeed != 5 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrais être à 5 et 0 mais il sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestMode2(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(0)
	if maxSpeed != 5 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrais être à 5 et 0 mais il sont à ", maxSpeed, " et ", minSpeed)
	}
}

func TestMode3(t *testing.T) {
	maxSpeed, minSpeed := MinAndMaxSpeed(0)
	if maxSpeed != 5 || minSpeed != 0 {
		t.Error("maxSpeed et minSpeed devrais être à 5 et 0 mais il sont à ", maxSpeed, " et ", minSpeed)
	}
}
