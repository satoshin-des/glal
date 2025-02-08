package mat

import "math"

func Trace(mat Matrix) float64 {
	var s float64 = 0

	for i := 0; i < int(math.Min(float64(mat.nrows), float64(mat.ncols))); i++ {
		s += mat.at[i][i]
	}

	return s
}
