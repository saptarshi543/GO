package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	// this function can take n no. of arguments...
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {
	fmt.Printf("%d", sum(1, 243, 53, 2, 3, 5, 25, 13))
}
