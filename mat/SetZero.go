package mat

// SetZero sets all elements of mat to 0
func SetZero(mat Matrix) {
	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			mat.At[i][j] = 0
		}
	}
}
