package mat

// SetOne sets all elements of mat to 1
func SetOne(mat Matrix) {
	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			mat.At[i][j] = 1
		}
	}
}
