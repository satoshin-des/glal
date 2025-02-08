package mat

// Trans returns transpose of mat
func Trans(mat Matrix) Matrix {
	m := SetDims(mat.ncols, mat.nrows)

	for i := 0; i < m.nrows; i++ {
		for j := 0; j < m.ncols; j++ {
			m.at[i][j] = mat.at[j][i]
		}
	}

	return m
}
