package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	fmt.Println(str)
}
