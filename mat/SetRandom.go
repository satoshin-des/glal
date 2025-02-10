package mat

import "math/rand"

// SetRandom sets mat to random matrix
func SetRandom(mat Matrix) {
	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			mat.At[i][j] = float64(rand.Int63n(100))
		}
	}
}
