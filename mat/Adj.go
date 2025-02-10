package mat

// Adj returns adjugate of mat
func Adj(mat Matrix) Matrix {
	if !IsSquare(mat) {
		panic("adjugate of non-square matrix is not defined")
	}

	tempMat := SetDims(mat.NumRows, mat.NumCols)

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			tempMat.At[j][i] = CoFactor(mat, i, j)
		}
	}
	return tempMat
}
