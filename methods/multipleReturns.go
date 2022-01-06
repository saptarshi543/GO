package main

import (
	"fmt"
)

func values() (int, int) {
	return 12, 45
}
func main() {
	x, y := values()
	fmt.Printf("x: %d\ny: %d", x, y)
}
