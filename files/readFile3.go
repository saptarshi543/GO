package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// working.....
// does the same thing but stores the lines in an array
func read_lines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	nrt := bufio.NewScanner(file)
	for nrt.Scan() {
		lines = append(lines, nrt.Text())
	}
	return lines, nrt.Err()
}
func main() {
	lines, err := read_lines("test.txt")
	if err != nil {
		log.Fatalf("read_lines %s", err)
	}
	for i, lines := range lines {
		fmt.Println(i, lines)
	}
}
