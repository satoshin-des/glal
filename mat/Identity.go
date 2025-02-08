package mat

// Identity returns n-th identity matrix
func Identity(n int) Matrix {
	mat := SetDims(n, n)
	SetIdentity(mat)
	return mat
}
