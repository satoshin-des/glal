package mat

// SetDims returns n * m zero matrix
//
// panic if n or m is not positive
func SetDims(n int, m int) Matrix {
	if n < 1 || m < 1 {
		panic("size of matrix must be positive")
	}

	var mat Matrix

	mat.nrows = n
	mat.ncols = m
	mat.at = make([][]float64, n)
	for i := 0; i < n; i++ {
		mat.at[i] = make([]float64, m)
	}

	return mat
}
