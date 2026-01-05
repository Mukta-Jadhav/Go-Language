package main
import("fmt")
func main(){
	var a int
	var b int
	var choice int
	fmt.Printf("Enter the key to operation \n 1.addition \n 2.substraction \n 3.Multiplication \n 4.Division\n")
    fmt.Scanf("%d",&choice)
	fmt.Print("enter value of a:")
	fmt.Scan(&a)
	fmt.Print("enter value of b:")
	fmt.Scan(&b)
	switch (choice){
	case 1:
		fmt.Println("addition of",a, "and",b,"is:",a+b)
	
	case 2:
		fmt.Println("substraction of",a ," and" , "is:",a-b)
	case 3:
		fmt.Println("multiplication of",a ," and" , "is:",a*b)
	case 4:
		fmt.Println("division of",a ," and" , "is:",a/b)	
	default:
		fmt.Println("error choice from above")	

		

	}

	}    