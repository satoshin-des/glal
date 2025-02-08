package vec

import "math"

// TriArea returns area of triangle composed of x, y, and z
func TriArea(x Vector, y Vector, z Vector) float64 {
	xy := Norm(Sub(x, y), 2)
	yz := Norm(Sub(y, z), 2)
	zx := Norm(Sub(z, x), 2)

	s := (xy + yz + zx) * 0.5
	return math.Sqrt(s * (s - xy) * (s - yz) * (s - zx))
}
