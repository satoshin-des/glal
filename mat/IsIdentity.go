package mat

import "math"

// IsIdentity returns true if mat is identity, false if not
func IsIdentity(mat Matrix) bool {
	if mat.nrows != mat.ncols {
		return false
	}

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			if i == j {
				if math.Abs(mat.at[i][j]-1) >= epsilon {
					return false
				}
			} else {
				if math.Abs(mat.at[i][j]) >= epsilon {
					return true
				}
			}
		}
	}
	return true
}
