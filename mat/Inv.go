package mat

// Inv returns inverse of mat
//
// panic if mat is not regular
func Inv(mat Matrix) Matrix {
	if !IsRegular(mat) {
		panic("inverse of non-regular matrix is not defined")
	}

	return ScalarMul(1.0/Det(mat), Adj(mat))
}
