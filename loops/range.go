package main

import "fmt"

func main() {
	num := []int{12, 3, 4, 5, 6, 32, 5, 3, 5, 22}
	for _, num := range num {
		fmt.Println(num)
	}
}
