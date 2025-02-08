package vec

var epsilon float64 = 1e-6

// Vector is vector type
type Vector struct {
	length int
	at     []float64
}
