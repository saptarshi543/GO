package main

import (
	"fmt"
	"io/ioutil"
)

// working...
func main() {
	fmt.Printf("enter directory address of file to read\n")
	var address string
	fmt.Scanf("%s", &address)
	b, err := ioutil.ReadFile(address)
	if err != nil {
		fmt.Println(err)
	}
	// str:=string(b)
	fmt.Println(string(b))
}
