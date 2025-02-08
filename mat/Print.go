package mat

import "fmt"

// Print prints mat
func Print(mat Matrix) {
	fmt.Printf("[\n")
	for i := 0; i < mat.nrows; i++ {
		fmt.Printf("[%g", mat.at[i][0])
		for j := 1; j < mat.ncols; j++ {
			fmt.Printf(", %g", mat.at[i][j])
		}
		fmt.Printf("],\n")
	}
	fmt.Printf("]\n")
}
