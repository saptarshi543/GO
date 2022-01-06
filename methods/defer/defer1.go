package main

import (
	"fmt"
)

func first() {
	fmt.Printf("first\n")
}
func second() {
	fmt.Printf("second\n")
}
func main() {
	defer first()
	second()
}
