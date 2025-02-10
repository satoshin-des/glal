package glal

import (
	"fmt"
	"glal/mat"
	"glal/vec"
)

// MulMatVec returns m * v
//
// panic if m * v is not defined
func MulMatVec(m mat.Matrix, v vec.Vector) vec.Vector {
	if m.NumCols != v.Length {
		pn := fmt.Sprintf("product of %d by %d-matrix and %d-th vector is not defined", m.NumRows, m.NumCols, v.Length)
		panic(pn)
	}

	x := vec.SetLength(m.NumRows)

	for i := 0; i < m.NumRows; i++ {
		x.At[i] = 0
		for j := 0; j < m.NumCols; j++ {
			x.At[i] += m.At[i][j] * v.At[j]
		}
	}

	return x
}
