package vec

func Inner(x Vector, y Vector, z Vector) Vector {
	xy := Norm(Sub(x, y), 2)
	yz := Norm(Sub(y, z), 2)
	zx := Norm(Sub(z, x), 2)

	vec := SetLength(x.Length)

	for i := 0; i < vec.Length; i++ {
		vec.At[i] = (yz*x.At[i] + zx*y.At[i] + xy*z.At[i]) / (xy + yz + zx)
	}

	return vec
}
