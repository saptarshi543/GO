package main

import (
	"fmt"
)

type Person struct {
	name string
	job  string
	age  int
}

func main() {
	var Person_A Person
	Person_A.name = "person A"
	Person_A.job = "teacher"
	Person_A.age = 24
	fmt.Printf("Person A name: %s job: %s age: %d", Person_A.name, Person_A.job, Person_A.age)
}
