package mat

// ZeroMat returns n*m-zero matrix
//
// panic if n or m are not positive
func ZeroMat(n int, m int) Matrix {
	if n <= 0 || m <= 0 {
		panic("matrix whose size is not positive is not defined")
	}

	mat := SetDims(n, m)
	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			mat.At[i][j] = 0
		}
	}

	return mat
}
