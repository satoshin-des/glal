package mat

// swap i-th row and j-th row of mat
//
// panic if i or j are out of range
func SwapRow(mat Matrix, i int, j int) {
	if i >= mat.nrows || j >= mat.nrows || i < 0 || j < 0 {
		panic("index out of range")
	}

	if i != j {
		var t float64
		for k := 0; k < mat.ncols; k++ {
			t = mat.at[i][k]
			mat.at[i][k] = mat.at[j][k]
			mat.at[j][k] = t
		}
	}
}
