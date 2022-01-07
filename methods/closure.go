// The closure method returns a number that increases each time its called
package main

import (
	"fmt"
)

func main() {
	x := 2
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
