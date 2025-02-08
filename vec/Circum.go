package vec

// Circum returns circumcenter of x, y, and z
func Circum(x Vector, y Vector, z Vector) Vector {
	if x.length != y.length || y.length != z.length || z.length != x.length {
		panic("circumcenter of different three vectors is not defined")
	}

	area := TriArea(x, y, z)

	xy := Norm(Sub(x, y), 2)
	yz := Norm(Sub(y, z), 2)
	zx := Norm(Sub(z, x), 2)

	sinX := (area + area) / (xy * zx)
	sinY := (area + area) / (xy * yz)
	sinZ := (area + area) / (yz * zx)
	cosX := InnerProduct(Sub(z, x), Sub(x, y)) / (zx * xy)
	cosY := InnerProduct(Sub(x, y), Sub(y, z)) / (xy * yz)
	cosZ := InnerProduct(Sub(y, z), Sub(z, x)) / (yz * zx)

	vec := SetLength(x.length)

	for i := 0; i < vec.length; i++ {
		vec.at[i] = (sinX*cosX*x.at[i] + sinY*cosY*y.at[i] + sinZ*cosZ*z.at[i]) / (sinX*cosX + sinY*cosY + sinZ*cosZ)
	}

	return vec
}
