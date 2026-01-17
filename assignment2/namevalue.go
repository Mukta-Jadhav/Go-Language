package main
import ("fmt")

func multi( a,b int )int{
	result:=a*b
	return result
}

func main(){
	fmt.Println("multiplication is:",multi(13,6))
}