package mat

// CoFactor returns (row, col)-cofactor of mat
func CoFactor(mat Matrix, row int, col int) float64 {
	tempMat := SetDims(mat.nrows-1, mat.ncols-1)

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			if i < row {
				if j < col {
					tempMat.at[i][j] = mat.at[i][j]
				} else if j > col {
					tempMat.at[i][j-1] = mat.at[i][j]
				}
			} else if i > row {
				if j < col {
					tempMat.at[i-1][j] = mat.at[i][j]
				} else if j > col {
					tempMat.at[i-1][j-1] = mat.at[i][j]
				}
			}
		}
	}

	tempDet := Det(tempMat)

	if (row+col)%2 == 1 {
		tempDet = -tempDet
	}

	return tempDet
}
