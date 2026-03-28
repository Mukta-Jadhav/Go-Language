package main

import (
	"fmt"
	
)

// Define structure
type Student struct {
	roll_no   int
	stud_name string
	mark1     float64
	mark2     float64
	mark3     float64
	total     float64
	average   float64
}

func main() {
	var n int

	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	// Create slice of students
	students := make([]Student, n)

	// Input student details
	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details for student", i+1)

		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].roll_no)

		fmt.Print("Name: ")
		fmt.Scan(&students[i].stud_name)

		fmt.Print("Mark 1: ")
		fmt.Scan(&students[i].mark1)

		fmt.Print("Mark 2: ")
		fmt.Scan(&students[i].mark2)

		fmt.Print("Mark 3: ")
		fmt.Scan(&students[i].mark3)

		// Calculate total and average
		students[i].total = students[i].mark1 + students[i].mark2 + students[i].mark3
		students[i].average = students[i].total / 3
	}

	// Display results
	fmt.Println("\n--- Student Details ---")
	for i := 0; i < n; i++ {
		fmt.Println("\nStudent", i+1)
		fmt.Println("Roll No:", students[i].roll_no)
		fmt.Println("Name:", students[i].stud_name)
		fmt.Println("Total Marks:", students[i].total)
		fmt.Println("Average Marks:", students[i].average)
	}
}