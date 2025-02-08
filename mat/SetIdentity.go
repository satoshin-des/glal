package mat

// SetIdentity sets mat to identity matrix
func SetIdentity(mat Matrix) {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			if i == j {
				mat.at[i][j] = 1
			} else {
				mat.at[i][j] = 0
			}
		}
	}
}
