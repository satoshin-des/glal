package mat

// Sub returns difference of mat1 and mat2
//
// panic if size of mat1 and mat2 is different
func Sub(mat1 Matrix, mat2 Matrix) Matrix {
	if mat1.nrows != mat2.nrows || mat1.ncols != mat2.ncols {
		panic("difference of two different size matrix is not defined")
	}

	mat := SetDims(mat1.nrows, mat1.ncols)

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = mat1.at[i][j] - mat2.at[i][j]
		}
	}

	return mat
}
