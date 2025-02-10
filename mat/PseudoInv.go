package mat

// PseudoInv returns pseudo-inverse of mat
func PseudoInv(mat Matrix) Matrix {
	return Mul(Inv(Mul(Trans(mat), mat)), Trans(mat))
}
