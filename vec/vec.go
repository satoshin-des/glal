package vec

import (
	"fmt"
	"math"
	"math/rand"
)

// Vector is vector type
type Vector struct {
	length int
	at     []float64
}

// Prints Vector
func Print(vec Vector) {
	fmt.Printf("%g", vec.at[0])
	for i := 1; i < vec.length; i++ {
		fmt.Printf(", %g", vec.at[i])
	}
	fmt.Printf("\n")
}

// Makes vector whose size is n
func SetLength(n int) Vector {
	var vec Vector

	vec.at = make([]float64, n)
	vec.length = n
	for i := 0; i < n; i++ {
		vec.at[i] = 0
	}

	return vec
}

func SetZero(vec Vector) {
	for i := 0; i < vec.length; i++ {
		vec.at[i] = 0
	}
}

func SetOne(vec Vector) {
	for i := 0; i < vec.length; i++ {
		vec.at[i] = 1
	}
}

func SetRandom(vec Vector) {
	for i := 0; i < vec.length; i++ {
		vec.at[i] = float64(rand.Int63n(100))
	}
}

func Add(vec1 Vector, vec2 Vector) Vector {
	if vec1.length != vec2.length {
		panic("sum of two different size vector is not defined")
	}

	vec := SetLength(vec1.length)

	for i := 0; i < vec1.length; i++ {
		vec.at[i] = vec1.at[i] + vec2.at[i]
	}

	return vec
}

func Sub(vec1 Vector, vec2 Vector) Vector {
	if vec1.length != vec2.length {
		panic("difference of two different size vector is not defined")
	}

	vec := SetLength(vec1.length)

	for i := 0; i < vec1.length; i++ {
		vec.at[i] = vec1.at[i] - vec2.at[i]
	}

	return vec
}

func ScalarMul(x float64, vec Vector) Vector {
	v := SetLength(vec.length)

	for i := 0; i < vec.length; i++ {
		v.at[i] = x * vec.at[i]
	}

	return v
}

// Computes inner products of two vectors
func InnerProduct(vec1 Vector, vec2 Vector) float64 {
	if vec1.length != vec2.length {
		panic("inner products of two different size vector is not defined")
	}

	var s float64 = 0

	for i := 0; i < vec1.length; i++ {
		s += vec1.at[i] * vec2.at[i]
	}

	return s
}

// Computes cross products of two vector
func CrossProduct(vec1 Vector, vec2 Vector) Vector {
	if vec1.length != 3 || vec2.length != 3 {
		panic("cross products of vectors whose size is not 3 is not defined")
	}

	vec := SetLength(3)

	vec.at[0] = vec1.at[1]*vec2.at[2] - vec2.at[1]*vec1.at[1]
	vec.at[1] = vec1.at[2]*vec2.at[0] - vec2.at[2]*vec1.at[0]
	vec.at[2] = vec1.at[0]*vec2.at[1] - vec2.at[0]*vec1.at[1]

	return vec
}

// Computes norm of a vector
func Norm(vec Vector) float64 {
	var s float64 = 0

	for i := 0; i < vec.length; i++ {
		s += vec.at[i] * vec.at[i]
	}

	return math.Sqrt(s)
}
