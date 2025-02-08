package vec

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
