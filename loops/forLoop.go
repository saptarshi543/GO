package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter no. of times the loop should run: ")
	var count int
	fmt.Scanf("%d", &count)
	for x := 0; x < count; x++ {
		fmt.Printf("%d\n", x)
	}
}
