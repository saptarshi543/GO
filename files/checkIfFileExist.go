package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("test.txt"); err == nil {
		fmt.Println("File Exist\n")
	} else {
		fmt.Println("File Does Not Exist\n")
	}
}
