package main

import (
	"fmt"
	"math/rand"
)

// working...
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
func main() {
	fmt.Printf("enter max value and min value: \n")
	var max int
	var min int
	fmt.Scanf("%d\n%d", &max, &min)
	ans := random(max, min)
	fmt.Printf("--> %d", ans)
}
