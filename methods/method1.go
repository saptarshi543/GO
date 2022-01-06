package main

import (
	"fmt"
)

func hello1(name string) string {
	return "hello1 " + name + "\n"
}
func main() {
	var name string = "sap"
	fmt.Printf(hello1(name))
	hello2(name)
}
func hello2(name string) {
	fmt.Printf("hello2 %s", name)
}
