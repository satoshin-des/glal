package mat

import "fmt"

// Print prints mat
func Print(mat Matrix) {
	fmt.Printf("[\n")
	for i := 0; i < mat.NumRows; i++ {
		fmt.Printf("[%g", mat.At[i][0])
		for j := 1; j < mat.NumCols; j++ {
			fmt.Printf(", %g", mat.At[i][j])
		}
		fmt.Printf("],\n")
	}
	fmt.Printf("]\n")
}
