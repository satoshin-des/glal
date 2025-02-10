package mat

// SetIdentity sets mat to identity matrix
func SetIdentity(mat Matrix) {
	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			if i == j {
				mat.At[i][j] = 1
			} else {
				mat.At[i][j] = 0
			}
		}
	}
}
