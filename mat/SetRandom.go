package mat

import "math/rand"

// SetRandom sets mat to random matrix
func SetRandom(mat Matrix) {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = float64(rand.Int63n(100))
		}
	}
}
