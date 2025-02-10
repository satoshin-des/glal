package vec

func SetOne(vec Vector) {
	for i := 0; i < vec.Length; i++ {
		vec.At[i] = 1
	}
}
