package main

import (
	"fmt"
)

func squareSum(n int, ch chan int) {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n = n / 10
	}
	ch <- sum
}

func cubeSum(n int, ch chan int) {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d * d
		n = n / 10
	}
	ch <- sum
}

func main() {
	number := 123

	sqCh := make(chan int)
	cbCh := make(chan int)

	go squareSum(number, sqCh)
	go cubeSum(number, cbCh)

	squares := <-sqCh
	cubes := <-cbCh

	fmt.Println("Sum of squares =", squares)
	fmt.Println("Sum of cubes =", cubes)
	fmt.Println("Final sum of squares and cubes =", squares+cubes)
}
