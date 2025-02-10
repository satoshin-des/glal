package mat

// swap i-th row and j-th row of mat
//
// panic if i or j are out of range
func SwapRow(mat Matrix, i int, j int) {
	if i >= mat.NumRows || j >= mat.NumRows || i < 0 || j < 0 {
		panic("index out of range")
	}

	if i != j {
		var t float64
		for k := 0; k < mat.NumCols; k++ {
			t = mat.At[i][k]
			mat.At[i][k] = mat.At[j][k]
			mat.At[j][k] = t
		}
	}
}
