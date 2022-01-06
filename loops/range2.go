package main

import (
	"fmt"
)

func main() {
	var num = []int64{10, 9, 8, 7, 6, 5}
	for index, element := range num {
		fmt.Printf("index: %d   element: %d\n", index, element)
	}
}
