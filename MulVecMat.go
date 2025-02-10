package glal

import (
	"glal/mat"
	"glal/vec"
)

func MulVecMat(v vec.Vector, m mat.Matrix) vec.Vector {
	x := vec.SetLength(m.NumCols)

	for i := 0; i < m.NumCols; i++ {
		x.At[i] = 0
		for j := 0; j < m.NumRows; j++ {
			x.At[i] += v.At[i] * m.At[j][i]
		}
	}

	return x
}
