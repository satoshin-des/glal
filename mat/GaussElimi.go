package mat

// Det returns deteminant of mat
func GaussElimi(mat Matrix) Matrix {
	copyMat := SetDims(mat.nrows, mat.ncols)
	index := -1
	var temp float64

	var min int = 0
	if mat.nrows < mat.ncols {
		min = mat.nrows
	} else {
		min = mat.ncols
	}

	for k := 0; k < min; k++ {
		for i := k; i < mat.nrows; i++ {
			if mat.at[i][k] != 0 {
				index = i
				break
			}
		}

		if index == -1 {
			continue
		} else if index != k {
			SwapRow(mat, k, index)
		}

		for i := 0; i < mat.nrows; i++ {
			for j := k; j < mat.ncols; j++ {
				if i == k {
					copyMat.at[i][j] = mat.at[i][j]
				} else {
					copyMat.at[i][j] = mat.at[i][j] - mat.at[k][j]*mat.at[i][k]/mat.at[k][k]
				}
			}
		}

		temp = copyMat.at[k][k]

		for i := 0; i < mat.ncols; i++ {
			copyMat.at[k][i] /= temp
		}
	}

	return copyMat
}
