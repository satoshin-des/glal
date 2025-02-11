package glal

import (
	"github.com/satoshin-des/glal/mat"

	"github.com/satoshin-des/glal/vec"
)

// ArgMat returns argumented matrix of m and v
func ArgMat(m mat.Matrix, v vec.Vector) mat.Matrix {
	argMat := mat.SetDims(m.NumRows, m.NumCols+1)

	for i := 0; i < m.NumRows; i++ {
		for j := 0; j < m.NumCols; j++ {
			argMat.At[i][j] = m.At[i][j]
		}
	}

	for i := 0; i < m.NumRows; i++ {
		argMat.At[i][m.NumRows] = v.At[i]
	}

	return argMat
}
