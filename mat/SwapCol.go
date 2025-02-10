package mat

// swap i-th column and j-th column of mat
//
// panic if i or j are out of range
func SwapCol(mat Matrix, i int, j int) {
	if i >= mat.NumCols || j >= mat.NumCols || i < 0 || j < 0 {
		panic("index out of range")
	}

	if i != j {
		var t float64
		for k := 0; k < mat.NumRows; k++ {
			t = mat.At[k][i]
			mat.At[k][i] = mat.At[k][j]
			mat.At[k][j] = t
		}
	}
}
