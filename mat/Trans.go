package mat

// Trans returns transpose of mat
func Trans(mat Matrix) Matrix {
	m := SetDims(mat.NumCols, mat.NumRows)

	for i := 0; i < m.NumRows; i++ {
		for j := 0; j < m.NumCols; j++ {
			m.At[i][j] = mat.At[j][i]
		}
	}

	return m
}
