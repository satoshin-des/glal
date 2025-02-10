package glal

import (
	"glal/mat"
	"glal/vec"
)

func Mul(x any, y any) vec.Vector {
	var v vec.Vector
	switch x := x.(type) {
	case mat.Matrix:
		switch y := y.(type) {
		case vec.Vector:
			v = MulMatVec(x, y)
		}
	}
	return v
}
