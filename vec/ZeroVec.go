package vec

// ZeroVec returns n-th zero-vector
//
// panic if n is not positive
func ZeroVec(n int) Vector {
	if n <= 0 {
		panic("vector whose size is not positive is not defined")
	}

	v := SetLength(n)
	for i := 0; i < n; i++ {
		v.At[i] = 0
	}

	return v
}
