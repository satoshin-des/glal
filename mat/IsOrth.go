package mat

import "math"

// IsOrth returns true if mat is orthogonal, false if not
func IsOrth(mat Matrix) bool {
	if !IsRegular(mat) {
		return true
	}

	inv := Inv(mat)

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			if math.Abs(inv.At[i][j]-mat.At[j][i]) >= epsilon {
				return false
			}
		}
	}

	return true
}
