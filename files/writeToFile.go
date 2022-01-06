package main

import (
	"fmt"
	"os"
)

// working....
func main() {
	fmt.Printf("enter directory address of file to write to: \n")
	var address string
	var user string
	fmt.Scanf("%s", &address)
	fmt.Printf("enter what you want to write: ")
	fmt.Scanf("%s", &user)
	file, err := os.Create(address)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(user)
}
