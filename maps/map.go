package main

import (
	"fmt"
)

func main() {
	elements := make(map[string]string)
	elements["O"] = "oxygen"
	elements["Ca"] = "calcium"
	elements["C"] = "carbon"
	fmt.Println(elements["C"])
}
