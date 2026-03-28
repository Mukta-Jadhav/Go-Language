package main

import (
	"fmt"
)

type student struct {
	s_name    string
	branch    string
	college   string
	year  int
}

func (s *student) show() {
	fmt.Printf("Author Name: %s", s.s_name)
	fmt.Printf("\nBranch: %s and passing year: %d",
		s.branch, s.year)
	fmt.Printf("\nCollege Name: %s", s.college)
	
}




func main() {

	values := student{
		s_name:    "Shirwaikar",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2026,
	
		
	}
	s1:=student(values)
	s1.show()
	
}
