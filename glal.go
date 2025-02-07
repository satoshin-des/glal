package glal

import (
	"errors"
	"fmt"
)

// Vector is vector type whose elements have float64 type
type Vector struct {
	length int
	at     []float64
}

var (
	ErrOutOfIndex    = errors.New("out of index")
	ErrDifferentSize = errors.New("two vectors size is different")
)

func OutOfIndex(vec Vector, index int) error {
	if index >= vec.length {
		return ErrOutOfIndex
	}
	return nil
}

// Prints Vector
func PrintVec(vec Vector) {
	fmt.Printf("%g", vec.at[0])
	for i := 1; i < vec.length; i++ {
		fmt.Printf(", %g", vec.at[i])
	}
	fmt.Printf("\n")
}

// Makes zero vector whose size is n
func SetZero(n int) Vector {
	var vec Vector

	vec.at = make([]float64, n)
	vec.length = n
	for i := 0; i < n; i++ {
		vec.at[i] = 0
	}

	return vec
}

// Computes inner products of two vectors
func InnerProduct(vec1 Vector, vec2 Vector) float64 {
	var s float64 = 0

	for i := 0; i < vec1.length; i++ {
		s += vec1.at[i] * vec2.at[i]
	}

	return s
}
