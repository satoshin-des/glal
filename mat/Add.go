package mat

// Add returns sum of mat1 and mat2
//
// panic if size of mat1 and mat2 is different
func Add(mat1 Matrix, mat2 Matrix) Matrix {
	if mat1.NumRows != mat2.NumRows || mat1.NumCols != mat2.NumCols {
		panic("sum of two different size matrix is not defined")
	}

	mat := SetDims(mat1.NumRows, mat1.NumCols)

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			mat.At[i][j] = mat1.At[i][j] + mat2.At[i][j]
		}
	}

	return mat
}
