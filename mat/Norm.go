package mat

import "math"

// Norm returns norm of mat
func Norm(mat Matrix, p int) float64 {
	var s float64 = 0

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			s += math.Pow(math.Abs(mat.At[i][j]), float64(p))
		}
	}

	return math.Pow(s, 1.0/float64(p))
}
