package mat

import "math"

// IsToeplitz returns true if mat is Toeplitz, false if not
func IsToeplitz(mat Matrix) bool {
	if !IsSquare(mat) {
		return false
	}

	var temp float64

	for i := 0; i < mat.NumRows; i++ {
		temp = mat.At[i][0]
		for j := 1; j < mat.NumCols-i; j++ {
			if math.Abs(temp-mat.At[i+j][j]) >= epsilon {
				return false
			}
		}

		temp = mat.At[0][i]
		for j := 1; j < mat.NumCols-i; j++ {
			if math.Abs(temp-mat.At[j][i+j]) >= epsilon {
				return false
			}
		}
	}

	return true
}
