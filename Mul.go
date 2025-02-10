package glal

import (
	"glal/mat"
	"glal/vec"
)

// product of vector and matrix
func Mul(x any, y any) vec.Vector {
	var v vec.Vector

	switch x := x.(type) {
	case mat.Matrix:
		switch y := y.(type) {
		case vec.Vector:
			v = MulMatVec(x, y)
		}
	case vec.Vector:
		switch y := y.(type) {
		case mat.Matrix:
			v = MulVecMat(x, y)
		}
	}

	return v
}
