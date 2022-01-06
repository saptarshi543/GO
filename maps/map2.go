package main

import (
	"fmt"
)

func main() {
	l := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for i, k := range l {
		fmt.Println(k, i)
	}
}
