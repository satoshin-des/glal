package glal

import (
	"fmt"
	"glal/mat"
	"glal/vec"
)

// MulVecMat returns v * m
//
// panic if v * m is not defined
func MulVecMat(v vec.Vector, m mat.Matrix) vec.Vector {
	if m.NumCols != v.Length {
		pn := fmt.Sprintf("product of  %d-th vector and %d by %d-matrix is not defined", v.Length, m.NumRows, m.NumCols)
		panic(pn)
	}

	x := vec.SetLength(m.NumCols)

	for i := 0; i < m.NumCols; i++ {
		x.At[i] = 0
		for j := 0; j < m.NumRows; j++ {
			x.At[i] += v.At[j] * m.At[j][i]
		}
	}

	return x
}
