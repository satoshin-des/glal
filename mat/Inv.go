package mat

// Inv returns inverse of mat
//
// panic if mat is not regular
func Inv(mat Matrix) Matrix {
	if !IsRegular(mat) {
		panic("inverse of non-regular matrix is not defined")
	}

	copyMat := CopyMat(mat)
	elimiMat := CopyMat(mat)
	InvMat := SetDims(mat.NumRows, mat.NumCols)
	SetIdentity(InvMat)

	index := -1
	var temp float64

	var min int = 0
	if copyMat.NumRows < copyMat.NumCols {
		min = copyMat.NumRows
	} else {
		min = copyMat.NumCols
	}

	for k := 0; k < min; k++ {
		for i := k; i < copyMat.NumRows; i++ {
			if copyMat.At[i][k] != 0 {
				index = i
				break
			}
		}

		if index == -1 {
			continue
		} else if index != k {
			SwapRow(copyMat, k, index)
			SwapRow(InvMat, k, index)
		}

		for i := 0; i < copyMat.NumRows; i++ {
			for j := 0; j < copyMat.NumCols; j++ {
				if i == k {
					elimiMat.At[i][j] = copyMat.At[i][j]
				} else {
					elimiMat.At[i][j] = copyMat.At[i][j] - copyMat.At[k][j]*copyMat.At[i][k]/copyMat.At[k][k]
					InvMat.At[i][j] = InvMat.At[i][j] - InvMat.At[k][j]*copyMat.At[i][k]/copyMat.At[k][k]
				}
			}
		}

		temp = elimiMat.At[k][k]

		for i := 0; i < copyMat.NumCols; i++ {
			elimiMat.At[k][i] /= temp
			InvMat.At[k][i] /= temp
		}

		copyMat = CopyMat(elimiMat)
	}

	return InvMat
}
