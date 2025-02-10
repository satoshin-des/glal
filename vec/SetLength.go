package vec

// Makes vector whose size is n
func SetLength(n int) Vector {
	var vec Vector

	vec.At = make([]float64, n)
	vec.Length = n
	for i := 0; i < n; i++ {
		vec.At[i] = 0
	}

	return vec
}
