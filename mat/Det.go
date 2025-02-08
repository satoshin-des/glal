package mat

// Det returns deteminant of mat
func Det(mat Matrix) float64 {
	if mat.ncols != mat.nrows {
		panic("determinant of non-square matrix is not defined")
	}

	var s float64 = 0
	m := SetDims(mat.nrows-1, mat.ncols-1)
	copyMat := SetDims(mat.nrows, mat.ncols)
	index := -1

	for i := 0; i < mat.ncols; i++ {
		if mat.at[0][i] != 0 {
			index = i
			break
		}
	}

	if index == -1 {
		return 0
	} else if index != 0 {
		SwapCol(mat, 0, index)
	}

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.nrows; j++ {
			if j == 0 {
				copyMat.at[i][j] = mat.at[i][j]
			} else {
				copyMat.at[i][j] = mat.at[i][j] - mat.at[i][0]*mat.at[0][j]/mat.at[0][0]
			}
		}
	}

	if copyMat.nrows == 2 {
		return copyMat.at[0][0] * copyMat.at[1][1]
	} else {
		for i := 1; i < mat.nrows; i++ {
			for j := 1; j < mat.ncols; j++ {
				m.at[i-1][j-1] = copyMat.at[i][j]
			}
		}
		s = copyMat.at[0][0] * Det(m)
	}
	return s
}
