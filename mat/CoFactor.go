package mat

// CoFactor returns (row, col)-cofactor of mat
func CoFactor(mat Matrix, row int, col int) float64 {
	tempMat := SetDims(mat.NumRows-1, mat.NumCols-1)

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			if i < row {
				if j < col {
					tempMat.At[i][j] = mat.At[i][j]
				} else if j > col {
					tempMat.At[i][j-1] = mat.At[i][j]
				}
			} else if i > row {
				if j < col {
					tempMat.At[i-1][j] = mat.At[i][j]
				} else if j > col {
					tempMat.At[i-1][j-1] = mat.At[i][j]
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
