package mat

// Det returns deteminant of mat
func GaussElimi(mat Matrix) Matrix {
	elimiMat := CopyMat(mat)
	copyMat := CopyMat(mat)

	index := -1
	var temp float64

	var min int = 0
	if copyMat.nrows < copyMat.ncols {
		min = copyMat.nrows
	} else {
		min = copyMat.ncols
	}

	for k := 0; k < min; k++ {
		for i := k; i < copyMat.nrows; i++ {
			if copyMat.at[i][k] != 0 {
				index = i
				break
			}
		}

		if index == -1 {
			continue
		} else if index != k {
			SwapRow(copyMat, k, index)
		}

		for i := 0; i < copyMat.nrows; i++ {
			for j := 0; j < copyMat.ncols; j++ {
				if i == k {
					elimiMat.at[i][j] = copyMat.at[i][j]
				} else {
					elimiMat.at[i][j] = copyMat.at[i][j] - copyMat.at[k][j]*copyMat.at[i][k]/copyMat.at[k][k]
				}
			}
		}

		temp = elimiMat.at[k][k]

		for i := 0; i < copyMat.ncols; i++ {
			elimiMat.at[k][i] /= temp
		}
	}

	return elimiMat
}
