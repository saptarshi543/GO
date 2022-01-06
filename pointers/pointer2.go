package main

import (
	"fmt"
)

func main() {
	var x int = 5
	var pointer *int
	pointer = &x
	fmt.Printf("Memory address: %x\n", &x)
	fmt.Printf("memory address in pointer: %x\n", pointer)
	fmt.Printf("value: %d", *pointer)
}
