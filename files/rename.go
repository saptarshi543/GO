package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("enter directory address of file to read\n")
	var address string
	var rename_to string
	fmt.Scanf("%s", &address)
	fmt.Printf("what do you want to rename it to?: ")
	fmt.Scanf("%s", &rename_to)
	os.Rename(address, rename_to)
}
