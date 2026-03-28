package main

import (
	"fmt"
)

type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

func (a author) show() {
	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d",
		a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}




func main() {

	values := author{
		a_name:    "Shirwaikar",
		branch:    "Computer science",
		college:   "XYZ",
		year:      1990,
		salary:    80000,
		
	}
	a1:=author(values)
	a1.show()
	
}
