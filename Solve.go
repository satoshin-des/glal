package glal

import (
	"errors"

	"github.com/satoshin-des/glal/mat"

	"github.com/satoshin-des/glal/vec"
)

// Solve returns one of slotion vectors of equation m * x = v
//
// returns pseudo solution and error if m * x = v has no solution
//
// panic if number of rows of m and length of v is defferent
func Solve(m mat.Matrix, v vec.Vector) (vec.Vector, error) {
	if m.NumRows != v.Length {
		panic("the given equation is not defined")
	}

	pseudoInv := mat.PseudoInv(m)

	if mat.Rank(m) != mat.Rank(ArgMat(m, v)) {
		return MulMatVec(pseudoInv, v), errors.New("given equation has not solution")
	} else {
		return MulMatVec(pseudoInv, v), nil
	}
}
