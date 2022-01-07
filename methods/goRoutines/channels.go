package main

import (
	"fmt"
)

func f1(c chan string) {
	c <- "f() was here"
}
func f2(c chan string) {
	msg := <-c
	fmt.Println("f2", msg)
}
func main() {
	var c chan string = make(chan string)
	go f1(c)
	go f2(c)
	fmt.Scanln()
}
