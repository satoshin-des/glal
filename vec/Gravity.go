package vec

// Gravity returns center of gravity of x, y, and z
//
// panic if size of three vectors x, y, and z are different
func Gravity(x Vector, y Vector, z Vector) Vector {
	if x.length != y.length || y.length != z.length || z.length != x.length {
		panic("center of gravity of different three vectors")
	}

	tempVec := Add(x, Add(y, z))
	return ScalarMul(1.0/3.0, tempVec)
}
