package mat

// IsSquare returns true if mat is square, false if not
func IsSquare(mat Matrix) bool {
	if mat.nrows == mat.ncols {
		return true
	} else {
		return false
	}
}
