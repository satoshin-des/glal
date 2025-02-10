package vec

func SetZero(vec Vector) {
	for i := 0; i < vec.Length; i++ {
		vec.At[i] = 0
	}
}
