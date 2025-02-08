package vec

import "fmt"

// Print prints vec
func Print(vec Vector) {
	fmt.Printf("[%g", vec.at[0])
	for i := 1; i < vec.length; i++ {
		fmt.Printf(", %g", vec.at[i])
	}
	fmt.Printf("]\n")
}
