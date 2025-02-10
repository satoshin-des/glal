package mat

import "math"

// IsSym returns true if mat is symmetric, false if not
func IsSym(mat Matrix) bool {
	if !IsSquare(mat) {
		return false
	}

	for i := 0; i < mat.NumRows; i++ {
		for j := i; j < mat.NumCols; j++ {
			if math.Abs(mat.At[i][j]-mat.At[j][i]) >= epsilon {
				return false
			}
		}
	}

	return true
}
