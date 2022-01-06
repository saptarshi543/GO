package main

import (
	"fmt"
)

func main() {
	switch a := 1; {
	case a == 1:
		fmt.Println("integer is 1")
		fallthrough
	case a == 2:
		fmt.Println("integer is 2")
	case a == 3:
		fmt.Println("integer is 3")
		fallthrough
	case a == 4:
		fmt.Println("integer is 4")
	default:
		fmt.Println("default..")
	}
}
