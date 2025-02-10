package mat

// CopyMat returns copy of mat
func CopyMat(mat Matrix) Matrix {
	m := SetDims(mat.NumRows, mat.NumCols)

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			m.At[i][j] = mat.At[i][j]
		}
	}

	return m
}
