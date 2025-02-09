package mat

// CopyMat returns copy of mat
func CopyMat(mat Matrix) Matrix {
	m := SetDims(mat.nrows, mat.ncols)

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			m.at[i][j] = mat.at[i][j]
		}
	}

	return m
}
