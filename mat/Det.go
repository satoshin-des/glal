package mat

// Det returns deteminant of mat
func Det(mat Matrix) float64 {
	if mat.NumCols != mat.NumRows {
		panic("determinant of non-square matrix is not defined")
	}

	var s float64 = 0
	m := SetDims(mat.NumRows-1, mat.NumCols-1)
	copyMat := SetDims(mat.NumRows, mat.NumCols)
	index := -1

	for i := 0; i < mat.NumCols; i++ {
		if mat.At[0][i] != 0 {
			index = i
			break
		}
	}

	if index == -1 {
		return 0
	} else if index != 0 {
		SwapCol(mat, 0, index)
	}

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumRows; j++ {
			if j == 0 {
				copyMat.At[i][j] = mat.At[i][j]
			} else {
				copyMat.At[i][j] = mat.At[i][j] - mat.At[i][0]*mat.At[0][j]/mat.At[0][0]
			}
		}
	}

	if copyMat.NumRows == 2 {
		return copyMat.At[0][0] * copyMat.At[1][1]
	} else {
		for i := 1; i < mat.NumRows; i++ {
			for j := 1; j < mat.NumCols; j++ {
				m.At[i-1][j-1] = copyMat.At[i][j]
			}
		}
		s = copyMat.At[0][0] * Det(m)
	}
	return s
}
