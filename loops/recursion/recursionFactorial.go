package main

import (
	"fmt"
)

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println("enter a number")
	var a int
	fmt.Scanf("%d", &a)
	fmt.Printf("factorial: %d", factorial(a))
}
