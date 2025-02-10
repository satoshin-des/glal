package vec

// Computes inner products of two vectors
func InnerProduct(vec1 Vector, vec2 Vector) float64 {
	if vec1.Length != vec2.Length {
		panic("inner products of two different size vector is not defined")
	}

	var s float64 = 0

	for i := 0; i < vec1.Length; i++ {
		s += vec1.At[i] * vec2.At[i]
	}

	return s
}
