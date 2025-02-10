package vec

import "math"

func IsZero(vec Vector) bool {
	for i := 0; i < vec.Length; i++ {
		if math.Abs(vec.At[i]) >= epsilon {
			return false
		}
	}
	return true
}
