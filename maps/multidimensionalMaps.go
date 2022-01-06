package main

import (
	"fmt"
)

func main() {
	collections := map[string]map[string]string{
		"Apple": map[string]string{
			"iphone":  "iphone 13 pro max",
			"tablets": "ipad pro 12 inches",
			"laptops": "macbook pro 16 inch",
		},
		"Microsoft": map[string]string{
			"os":                     "windows 11",
			"vs code":                "visual studio code 2022",
			"language(made by them)": "c# or cSharp",
		},
	}
	fmt.Println(collections["Apple"]["iphone"])
	fmt.Println(collections["Apple"]["laptops"])
	fmt.Println(collections["Microsoft"]["os"])
}
