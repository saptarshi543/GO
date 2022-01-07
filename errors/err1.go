package main

import (
	"errors"
	"fmt"
)

func do() (int, error) {
	return -1, errors.New("something went wrong !!")
}
func main() {
	fmt.Println(do())
}
