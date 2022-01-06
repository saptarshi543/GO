package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
