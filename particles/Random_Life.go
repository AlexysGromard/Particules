package particles

import (
	//"container/list"

	"math/rand"
	//"fmt"
)

func Random_Life() (life int) {
	life = rand.Intn(50)
	return life
}
