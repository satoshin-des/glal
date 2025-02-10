package mat

func ScalarMul(x float64, mat Matrix) Matrix {
	m := SetDims(mat.NumRows, mat.NumCols)

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			m.At[i][j] = x * mat.At[i][j]
		}
	}

	return m
}
