package vec

func ScalarMul(x float64, vec Vector) Vector {
	v := SetLength(vec.Length)

	for i := 0; i < vec.Length; i++ {
		v.At[i] = x * vec.At[i]
	}

	return v
}
