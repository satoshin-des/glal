package mat

func GramSchmidt(mat Matrix) Matrix {
	gsoMat := SetDims(mat.NumRows, mat.NumCols)
	mu := Identity(mat.NumCols)

	var dotTemp, normTemp float64

	for i := 0; i < mat.NumCols; i++ {
		mu.At[i][i] = 1

		for j := 0; j < mat.NumRows; j++ {
			gsoMat.At[j][i] = mat.At[j][i]
		}

		for j := 0; j < i; j++ {
			dotTemp = 0
			for k := 0; k < mat.NumRows; k++ {
				dotTemp += mat.At[k][i] * mat.At[k][j]
				normTemp += gsoMat.At[k][j] * gsoMat.At[k][j]
			}

			mu.At[j][i] = dotTemp / normTemp

			for k := 0; k < mat.NumRows; k++ {
				gsoMat.At[k][i] -= mu.At[j][i] * gsoMat.At[k][j]
			}
		}
	}

	Print(Mul(gsoMat, mu))

	return gsoMat
}
