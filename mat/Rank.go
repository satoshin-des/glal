package mat

import "math"

// Rank returns rank of mat
func Rank(mat Matrix) int {
	elimiMat := GaussElimi(mat)
	var rank int = 0

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			if math.Abs(elimiMat.At[i][j]) < epsilon {
				rank++
				break
			}
		}
	}

	return rank
}
