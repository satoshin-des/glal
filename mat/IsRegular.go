package mat

import "math"

// IsRegular returns true if mat is regular, false if mat is not
func IsRegular(mat Matrix) bool {
	if mat.NumRows != mat.NumCols {
		return false
	}

	if math.Abs(Det(mat)) < epsilon {
		return false
	} else {
		return true
	}
}
