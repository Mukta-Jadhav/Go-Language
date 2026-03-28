package main

import "fmt"

func main() {
	a := [][]int{{10, 20, 30, 40, 50}, {60, 70, 80, 90, 100}}
	fmt.Println("a=", a)
	fmt.Println("Accesing element from a as matrix:\n")
	for i := 0; i < len(a); i++ {

		fmt.Println(a[i])
	}
}