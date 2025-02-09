package mat

import "math"

// Rank returns rank of mat
func Rank(mat Matrix) int {
	elimiMat := GaussElimi(mat)
	var rank int = 0

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			if math.Abs(elimiMat.at[i][j]) < epsilon {
				rank++
				break
			}
		}
	}

	return rank
}
