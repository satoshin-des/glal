package vec

import "fmt"

// Print prints vec
func Print(vec Vector) {
	fmt.Printf("[%g", vec.At[0])
	for i := 1; i < vec.Length; i++ {
		fmt.Printf(", %g", vec.At[i])
	}
	fmt.Printf("]\n")
}
