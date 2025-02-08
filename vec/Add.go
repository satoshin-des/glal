package vec

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
