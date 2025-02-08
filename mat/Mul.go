package mat

// Mul returns product of mat1 and mat2
//
// panic if product of mat1 and mat2 is not defined
func Mul(mat1 Matrix, mat2 Matrix) Matrix {
	if mat1.ncols != mat2.nrows {
		panic("products of given matrix is not defined")
	}

	mat := SetDims(mat1.nrows, mat2.ncols)

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = 0
			for k := 0; k < mat1.ncols; k++ {
				mat.at[i][j] += mat1.at[i][k] * mat2.at[k][j]
			}
		}
	}

	return mat
}
