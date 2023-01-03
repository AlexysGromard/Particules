package particles

func MinAndMaxSpeed(SpeedMode int) (minSpeed, maxSpeed int) {
	switch SpeedMode {
	case 0:
		minSpeed = 0
		maxSpeed = 5
	case 1:
		minSpeed = 0
		maxSpeed = 1
	case 2:
		minSpeed = 1
		maxSpeed = 3
	case 3:
		minSpeed = 3
		maxSpeed = 5
	}

	return maxSpeed, minSpeed
}
