package mat

// SetOne sets all elements of mat to 1
func SetOne(mat Matrix) {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = 1
		}
	}
}
