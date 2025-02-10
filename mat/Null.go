package mat

// Null returns nullity of mat
func Null(mat Matrix) int {
	return mat.NumRows - Rank(mat)
}
