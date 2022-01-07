package main

import (
	"fmt"
)

func main() {
	var c chan string = make(chan string, 3)
	c <- "hello"
	c <- "world"
	c <- "golang"
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
