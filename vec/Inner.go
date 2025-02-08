package vec

func Inner(x Vector, y Vector, z Vector) Vector {
	xy := Norm(Sub(x, y), 2)
	yz := Norm(Sub(y, z), 2)
	zx := Norm(Sub(z, x), 2)

	vec := SetLength(x.length)

	for i := 0; i < vec.length; i++ {
		vec.at[i] = (yz*x.at[i] + zx*y.at[i] + xy*z.at[i]) / (xy + yz + zx)
	}

	return vec
}
