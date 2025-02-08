package mat

var epsilon float64 = 1e-6

// Matrix is matrix type
type Matrix struct {
	nrows int
	ncols int
	at    [][]float64
}
