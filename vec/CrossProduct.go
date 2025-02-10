package vec

// Computes cross products of two vector
func CrossProduct(vec1 Vector, vec2 Vector) Vector {
	if vec1.Length != 3 || vec2.Length != 3 {
		panic("cross products of vectors whose size is not 3 is not defined")
	}

	vec := SetLength(3)

	vec.At[0] = vec1.At[1]*vec2.At[2] - vec2.At[1]*vec1.At[1]
	vec.At[1] = vec1.At[2]*vec2.At[0] - vec2.At[2]*vec1.At[0]
	vec.At[2] = vec1.At[0]*vec2.At[1] - vec2.At[0]*vec1.At[1]

	return vec
}
