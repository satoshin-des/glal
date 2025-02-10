package vec

// Circum returns circumcenter of x, y, and z
func Circum(x Vector, y Vector, z Vector) Vector {
	if x.Length != y.Length || y.Length != z.Length || z.Length != x.Length {
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

	vec := SetLength(x.Length)

	for i := 0; i < vec.Length; i++ {
		vec.At[i] = (sinX*cosX*x.At[i] + sinY*cosY*y.At[i] + sinZ*cosZ*z.At[i]) / (sinX*cosX + sinY*cosY + sinZ*cosZ)
	}

	return vec
}
