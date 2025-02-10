package mat

import "math"

// IsIdentity returns true if mat is identity, false if not
func IsIdentity(mat Matrix) bool {
	if mat.NumRows != mat.NumCols {
		return false
	}

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			if i == j {
				if math.Abs(mat.At[i][j]-1) >= epsilon {
					return false
				}
			} else {
				if math.Abs(mat.At[i][j]) >= epsilon {
					return true
				}
			}
		}
	}
	return true
}
