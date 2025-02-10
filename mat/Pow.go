package mat

// Pow returns n-th power of mat
func Pow(mat Matrix, n int) Matrix {
	tempMat := Identity(mat.NumRows)

	if n > 0 {
		for i := 0; i < mat.NumRows; i++ {
			tempMat = Mul(tempMat, mat)
		}
	} else if n < 0 {
		for i := 0; i < mat.NumRows; i++ {
			tempMat = Mul(tempMat, mat)
		}
		tempMat = Inv(tempMat)
	}
	return tempMat
}
