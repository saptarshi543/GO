package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(3 * time.Second)
	<-t1.C
	fmt.Println("Timer expired")
}
