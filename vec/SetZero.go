package vec

func SetZero(vec Vector) {
	for i := 0; i < vec.length; i++ {
		vec.at[i] = 0
	}
}
