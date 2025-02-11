package mat

import "github.com/satoshin-des/glal/vec"

// Diag returns matrix whose diagonal elements are v
func Diag(v vec.Vector) Matrix {
	mat := ZeroMat(v.Length, v.Length)

	for i := 0; i < mat.NumRows; i++ {
		mat.At[i][i] = v.At[i]
	}

	return mat
}
