package main

import (
	"fmt"
	"os"
)

// working...
func main() {
	fmt.Println("enter directory of file:\n")
	var address string
	fmt.Scanf("%s", &address)
	if _, err := os.Stat(address); err == nil {
		fmt.Println("File Exist\n")
	} else {
		fmt.Println("File Does Not Exist\n")
	}
}
