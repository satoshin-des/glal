package mat

// swap i-th column and j-th column of mat
//
// panic if i or j are out of range
func SwapCol(mat Matrix, i int, j int) {
	if i >= mat.ncols || j >= mat.ncols || i < 0 || j < 0 {
		panic("index out of range")
	}

	if i != j {
		var t float64
		for k := 0; k < mat.nrows; k++ {
			t = mat.at[k][i]
			mat.at[k][i] = mat.at[k][j]
			mat.at[k][j] = t
		}
	}
}
