package main

import "fmt"

// send only
func f(c chan<- string) {
	c <- "send only channels"
}

func main() {
	c := make(chan string, 1)
	f(c)
	fmt.Println(<-c)
}
