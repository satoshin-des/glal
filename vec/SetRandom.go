package vec

import "math/rand"

func SetRandom(vec Vector) {
	for i := 0; i < vec.Length; i++ {
		vec.At[i] = float64(rand.Int63n(100))
	}
}
