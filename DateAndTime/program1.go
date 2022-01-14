package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now().UTC()
	fmt.Println("Date: " + current.Format("2006 Jan 02"))
	fmt.Println("Time: " + current.Format("03:04:05"))
}
