package main

import (
	"fmt"
)

// Define a struct for Book
type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

func main() {
	var n int

	// Input number of books
	fmt.Print("Enter number of books: ")
	fmt.Scan(&n)

	// Create a slice of Book
	books := make([]Book, n)

	// Input book details
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Book %d:\n", i+1)

		fmt.Print("Book ID: ")
		fmt.Scan(&books[i].BookID)

		fmt.Print("Title: ")
		fmt.Scan(&books[i].Title)

		fmt.Print("Author: ")
		fmt.Scan(&books[i].Author)

		fmt.Print("Price: ")
		fmt.Scan(&books[i].Price)
	}

	// Display book details
	fmt.Println("\n--- Book Details ---")
	for i := 0; i < n; i++ {
		fmt.Printf("\nBook %d:\n", i+1)
		fmt.Println("Book ID:", books[i].BookID)
		fmt.Println("Title:", books[i].Title)
		fmt.Println("Author:", books[i].Author)
		fmt.Println("Price:", books[i].Price)
	}
}
