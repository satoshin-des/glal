package vec

import "math/rand"

func SetRandom(vec Vector) {
	for i := 0; i < vec.length; i++ {
		vec.at[i] = float64(rand.Int63n(100))
	}
}
