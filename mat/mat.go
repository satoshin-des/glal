package mat

var epsilon float64 = 1e-6

// Matrix is matrix type
type Matrix struct {
	NumRows int
	NumCols int
	At      [][]float64
}
