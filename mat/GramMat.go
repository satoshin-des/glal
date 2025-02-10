package mat

// GramMat returns Gram matrix of column vectors of mat
func GramMat(mat Matrix) Matrix {
	return Mul(Trans(mat), mat)
}
