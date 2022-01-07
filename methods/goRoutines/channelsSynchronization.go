package main

import (
	"fmt"
	"time"
)

func task(done chan bool) {
	fmt.Println("running")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
func main() {
	done := make(chan bool, 1)
	go task(done)
	<-done
}
