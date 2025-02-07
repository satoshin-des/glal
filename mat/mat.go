package mat

import (
	"fmt"
	"math"
	"math/rand"
)

// Matrix is matrix type
type Matrix struct {
	nrows int
	ncols int
	at    [][]float64
}

// Print prints mat
func Print(mat Matrix) {
	fmt.Printf("[\n")
	for i := 0; i < mat.nrows; i++ {
		fmt.Printf("[%g", mat.at[i][0])
		for j := 1; j < mat.ncols; j++ {
			fmt.Printf(", %g", mat.at[i][j])
		}
		fmt.Printf("],\n")
	}
	fmt.Printf("]\n")
}

// SetDims returns n * m zero matrix
func SetDims(n int, m int) Matrix {
	var mat Matrix

	mat.nrows = n
	mat.ncols = m
	mat.at = make([][]float64, n)
	for i := 0; i < n; i++ {
		mat.at[i] = make([]float64, m)
	}

	return mat
}

// SetZero sets all elements of mat to 0
func SetZero(mat Matrix) {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = 0
		}
	}
}

// SetIdentity sets mat to identity matrix
func SetIdentity(mat Matrix) {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			if i == j {
				mat.at[i][j] = 1
			} else {
				mat.at[i][j] = 0
			}
		}
	}
}

// SetOne sets all elements of mat to 1
func SetOne(mat Matrix) {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = 1
		}
	}
}

// SetRandom sets mat to random matrix
func SetRandom(mat Matrix) {
	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = float64(rand.Int63n(100))
		}
	}
}

// Add returns sum of mat1 and mat2
//
// panic if size of mat1 and mat2 is different
func Add(mat1 Matrix, mat2 Matrix) Matrix {
	if mat1.nrows != mat2.nrows || mat1.ncols != mat2.ncols {
		panic("sum of two different size matrix is not defined")
	}

	mat := SetDims(mat1.nrows, mat1.ncols)

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			mat.at[i][j] = mat1.at[i][j] + mat2.at[i][j]
		}
	}

	return mat
}

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

func ScalarMul(x float64, mat Matrix) Matrix {
	m := SetDims(mat.nrows, mat.ncols)

	for i := 0; i < mat.nrows; i++ {
		for j := 0; j < mat.ncols; j++ {
			m.at[i][j] = x * mat.at[i][j]
		}
	}

	return m
}

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

// Trans returns transpose of mat
func Trans(mat Matrix) Matrix {
	m := SetDims(mat.ncols, mat.nrows)

	for i := 0; i < m.nrows; i++ {
		for j := 0; j < m.ncols; j++ {
			m.at[i][j] = mat.at[j][i]
		}
	}

	return m
}

func Trace(mat Matrix) float64 {
	var s float64 = 0

	for i := 0; i < int(math.Min(float64(mat.nrows), float64(mat.ncols))); i++ {
		s += mat.at[i][i]
	}

	return s
}

// Det returns determinant of mat
//
// panic if mat is not square
func Det(mat Matrix) float64 {
	if mat.ncols != mat.nrows {
		panic("determinant of non-square matrix is not defined")
	}

	var s float64 = 0
	m := SetDims(mat.nrows-1, mat.ncols-1)

	if mat.nrows == 2 {
		return mat.at[0][0]*mat.at[1][1] - mat.at[0][1]*mat.at[1][0]
	} else {
		for i := 0; i < mat.ncols; i++ {
			for j := 1; j < mat.nrows; j++ {
				for k := 0; k < mat.ncols; k++ {
					if k < i {
						m.at[j-1][k] = mat.at[j][k]
					}
					if k > i {
						m.at[j-1][k-1] = mat.at[j][k]
					}
				}
			}

			if i%2 == 0 {
				s += mat.at[0][i] * Det(m)
			} else {
				s -= mat.at[0][i] * Det(m)
			}
		}
	}

	return s
}
