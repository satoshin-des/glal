package mat

import "math"

// IsZero returns true if mat is zero, false if not
func IsZero(mat Matrix) bool {
	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			if math.Abs(mat.At[i][j]) >= epsilon {
				return false
			}
		}
	}
	return true
}
