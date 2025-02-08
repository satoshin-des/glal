package vec

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
