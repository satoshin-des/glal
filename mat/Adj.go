package mat

// Adj returns adjugate of mat
func Adj(mat Matrix) Matrix {
	if !IsSquare(mat) {
		panic("adjugate of non-square matrix is not defined")
	}

	tempMat := SetDims(mat.nrows, mat.ncols)

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			tempMat.at[j][i] = CoFactor(mat, i, j)
		}
	}
	return tempMat
}
