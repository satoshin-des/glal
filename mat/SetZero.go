package mat

// SetZero sets all elements of mat to 0
func SetZero(mat Matrix) {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = 0
		}
	}
}
