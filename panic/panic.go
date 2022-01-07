package main

import (
	"fmt"
)

// panic method takes 1 argument : a message to show
// panic stops the execution of the program when encountered
func main() {
	panic("something went wrong")
	fmt.Printf("golang")
}
