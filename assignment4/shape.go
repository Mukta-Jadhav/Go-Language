package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle structure
type Circle struct {
	radius float64
}

// Rectangle structure
type Rectangle struct {
	length float64
	width  float64
}

// Implement methods for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Implement methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	var r float64
	var l, w float64

	// Input for circle
	fmt.Print("Enter radius of circle: ")
	fmt.Scan(&r)

	c := Circle{radius: r}

	// Input for rectangle
	fmt.Print("Enter length and width of rectangle: ")
	fmt.Scan(&l, &w)

	rect := Rectangle{length: l, width: w}

	// Using interface
	var s Shape

	// Circle
	s = c
	fmt.Println("\nCircle:")
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	// Rectangle
	s = rect
	fmt.Println("\nRectangle:")
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}