package main

import (
	"fmt"
)

func main() {
	var r1, c1, r2, c2 int

	// Input dimensions
	fmt.Print("Enter rows and columns of first matrix: ")
	fmt.Scan(&r1, &c1)

	fmt.Print("Enter rows and columns of second matrix: ")
	fmt.Scan(&r2, &c2)

	// Check multiplication condition
	if c1 != r2 {
		fmt.Println("Matrix multiplication not possible!")
		return
	}

	// Declare matrices
	mat1 := make([][]int, r1)
	mat2 := make([][]int, r2)
	result := make([][]int, r1)

	for i := 0; i < r1; i++ {
		mat1[i] = make([]int, c1)
	}
	for i := 0; i < r2; i++ {
		mat2[i] = make([]int, c2)
	}
	for i := 0; i < r1; i++ {
		result[i] = make([]int, c2)
	}

	// Input first matrix
	fmt.Println("\nEnter elements of first matrix:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			fmt.Scan(&mat1[i][j])
		}
	}

	// Input second matrix
	fmt.Println("\nEnter elements of second matrix:")
	for i := 0; i < r2; i++ {
		for j := 0; j < c2; j++ {
			fmt.Scan(&mat2[i][j])
		}
	}

	// Matrix multiplication
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			for k := 0; k < c1; k++ {
				result[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	// Display result
	fmt.Println("\nResultant Matrix:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			fmt.Print(result[i][j], " ")
		}
		fmt.Println()
	}
}