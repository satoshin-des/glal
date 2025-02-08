package vec

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
