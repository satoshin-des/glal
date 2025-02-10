package vec

func Sub(vec1 Vector, vec2 Vector) Vector {
	if vec1.Length != vec2.Length {
		panic("difference of two different size vector is not defined")
	}

	vec := SetLength(vec1.Length)

	for i := 0; i < vec1.Length; i++ {
		vec.At[i] = vec1.At[i] - vec2.At[i]
	}

	return vec
}
