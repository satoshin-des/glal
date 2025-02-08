package mat

import "math"

// IsZero returns true if mat is zero, false if not
func IsZero(mat Matrix) bool {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			if math.Abs(mat.at[i][j]) >= epsilon {
				return false
			}
		}
	}
	return true
}
