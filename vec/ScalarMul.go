package vec

func ScalarMul(x float64, vec Vector) Vector {
	v := SetLength(vec.length)

	for i := 0; i < vec.length; i++ {
		v.at[i] = x * vec.at[i]
	}

	return v
}
