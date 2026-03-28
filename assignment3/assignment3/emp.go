package main

import (
	"fmt"
)

// Define structure
type Employee struct {
	eno    int
	ename  string
	salary float64
}

func main() {
	var n int

	fmt.Print("Enter number of employees: ")
	fmt.Scan(&n)

	// Create slice of employees
	emp := make([]Employee, n)

	// Input employee details
	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details for employee", i+1)

		fmt.Print("Employee No: ")
		fmt.Scan(&emp[i].eno)

		fmt.Print("Employee Name: ")
		fmt.Scan(&emp[i].ename)

		fmt.Print("Salary: ")
		fmt.Scan(&emp[i].salary)
	}

	// Find maximum salary
	maxSalary := emp[0].salary
	for i := 1; i < n; i++ {
		if emp[i].salary > maxSalary {
			maxSalary = emp[i].salary
		}
	}

	// Display employees with maximum salary
	fmt.Println("\nEmployees with Maximum Salary:", maxSalary)
	for i := 0; i < n; i++ {
		if emp[i].salary == maxSalary {
			fmt.Println("\nEmployee No:", emp[i].eno)
			fmt.Println("Employee Name:", emp[i].ename)
			fmt.Println("Salary:", emp[i].salary)
		}
	}
}