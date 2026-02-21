package main

import "fmt"

func main() {
	subject := []string{
		"c++",
		"java",
		"php",
		"html",
		"css",
	}
	fmt.Println("element in subject are:")
	for i, subject := range subject {
		fmt.Printf("index %d:%s\n", i, subject)

	}
}
