package mat

import "math"

func Trace(mat Matrix) float64 {
	var s float64 = 0

	for i := 0; i < int(math.Min(float64(mat.NumRows), float64(mat.NumCols))); i++ {
		s += mat.At[i][i]
	}

	return s
}
