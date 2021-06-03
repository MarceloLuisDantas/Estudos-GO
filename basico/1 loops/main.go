// lOOPS EM GO

package main

import (
	"fmt"
)

func main() {
	// While
	println("\nWhile ------------------")
	var sum1 int = 0
	for sum1 < 21 {
		fmt.Printf("%d, ", sum1)
		sum1++
	}

	// For C-Like
	println("\n\nFor C-Like -------------")
	var sum2 int = 0
	for i := 0; i < 21; i++ {
		fmt.Printf("%d, ", sum2)
		sum2++
	}

	// ForEver
	println("\n\nForEver ----------------")
	var sum3 int = 0
	for {
		fmt.Printf("%d, ", sum3)
		if sum3++; sum3 == 21 {
			break
		}
	}

	// For Range
	println("\n\nFor Range --------------")
	var lista = []int{1, 2, 3, 4, 5}
	for i, num := range lista {
		fmt.Printf("%d = %d\n", i, num)
	}
}
