package mat

// GramDet returns Gram determinant of column vectors of mat
func GramDet(mat Matrix) float64 {
	return Det(Mul(Trans(mat), mat))
}
