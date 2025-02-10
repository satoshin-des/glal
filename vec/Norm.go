package vec

import "math"

// Computes norm of a vector
func Norm(vec Vector, p int) float64 {
	var s float64 = 0

	for i := 0; i < vec.Length; i++ {
		s += vec.At[i] * vec.At[i]
	}

	return math.Pow(s, 1.0/float64(p))
}
