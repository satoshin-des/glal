package vec

import "math"

func IsZero(vec Vector) bool {
	for i := 0; i < vec.length; i++ {
		if math.Abs(vec.at[i]) >= epsilon {
			return false
		}
	}
	return true
}
