package main

import (
	"fmt"
)

func main() {
	var x int = 4
	fmt.Printf("Address: %x\n", &x)
	fmt.Printf("Value: %d\n", x)
}
