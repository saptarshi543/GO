// technically GO does not have while
// but it can be done in a for loop
package main

import (
	"fmt"
)

func main() {
	i := 1
	max := 11
	for i <= max {
		fmt.Println(i)
		i += 1
	}
}
