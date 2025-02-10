package mat

// IsSquare returns true if mat is square, false if not
func IsSquare(mat Matrix) bool {
	if mat.NumRows == mat.NumCols {
		return true
	} else {
		return false
	}
}
