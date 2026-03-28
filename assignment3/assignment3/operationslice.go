package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println("a=", a)
	fmt.Println("Accesing element from a as matrix:\n")
	for i := 0; i < len(a); i++ {

		fmt.Println(a[i])
	}

	b := []int{45, 76, 30, 25, 5}
	fmt.Println("b=", b)
	fmt.Println("Accesing element from b as matrix:\n")
	for i := 0; i < len(b); i++ {

		fmt.Println(b[i])
	}

	//append
	slice1 := append(a, 45, 10)
	fmt.Println("after append data slice a is is :", slice1)
	slice2 := append(a, b...)
	fmt.Println("after append data slice a is is :", slice2)

	//sort()
	fmt.Println("slice1:", slice1)
	sort.Ints(slice1)
	fmt.Println("slice1:", slice1)

}