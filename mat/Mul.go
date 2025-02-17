package mat

import "fmt"

// Mul returns product of mat1 and mat2
//
// panic if product of mat1 and mat2 is not defined
func Mul(mat1 Matrix, mat2 Matrix) Matrix {
	if mat1.NumCols != mat2.NumRows {
		pn := fmt.Sprintf("product of %d by %d matrix and %d by %d matrix is not defined", mat1.NumRows, mat1.NumCols, mat2.NumRows, mat2.NumCols)
		panic(pn)
	}

	mat := SetDims(mat1.NumRows, mat2.NumCols)

	for i := 0; i < mat.NumRows; i++ {
		for j := 0; j < mat.NumCols; j++ {
			mat.At[i][j] = 0
			for k := 0; k < mat1.NumCols; k++ {
				mat.At[i][j] += mat1.At[i][k] * mat2.At[k][j]
			}
		}
	}

	return mat
}
