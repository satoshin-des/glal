package mat

import "math"

// IsSym returns true if mat is symmetric, false if not
func IsSym(mat Matrix) bool {
	if !IsSquare(mat) {
		return false
	}

	for i := 0; i < mat.nrows; i++ {
		for j := i; j < mat.ncols; j++ {
			if math.Abs(mat.at[i][j]-mat.at[j][i]) >= epsilon {
				return false
			}
		}
	}

	return true
}
