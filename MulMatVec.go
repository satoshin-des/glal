package glal

import (
	"glal/mat"
	"glal/vec"
)

// MulMatVec returns m * v
func MulMatVec(m mat.Matrix, v vec.Vector) vec.Vector {
	x := vec.SetLength(m.NumRows)

	for i := 0; i < m.NumRows; i++ {
		x.At[i] = 0
		for j := 0; j < m.NumCols; j++ {
			x.At[i] += m.At[i][j] * v.At[j]
		}
	}

	return x
}
