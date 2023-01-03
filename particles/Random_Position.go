package particles

import (
	//"container/list"

	"math/rand"
	//"fmt"
)

func Random_Position(intervalX, intervalY int) (PositionX, PositionY float64) {
	PositionX = float64(rand.Intn(intervalX)) + rand.Float64()
	PositionY = float64(rand.Intn(intervalY)) + rand.Float64()

	return PositionX, PositionY
}
