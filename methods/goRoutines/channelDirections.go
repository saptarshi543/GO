package main

import "fmt"

//read only
func f(c <-chan string) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan string, 1)
	c <- "hello"
	f(c)
}
