package main

import (
	"fmt"
)

func main() {
	fmt.Println("enter your age: ")
	var age int
	fmt.Scanf("%d", &age)
	if age >= 18 {
		fmt.Println("you can vote")
	} else {
		fmt.Println("you can NOT vote")
	}
}
