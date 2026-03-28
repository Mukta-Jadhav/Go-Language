package main
import ("fmt")
type Addition int
func (a Addition) tentimes() int{
	return int (a+10)
}

func (a *Addition) fivetimes() int{
	return int (*a+5)
}

func main(){
	var num int
	fmt.Println("Enter positive number:")
	fmt.Scanf("%d",&num)
	n1:=Addition(num)
	fmt.Println("ten times of given n1 number is:",n1.tentimes())
	n2:=&n1
	fmt.Println("ten times of given n2 number is:",n2.tentimes())
	fmt.Println("five times of given n2 number is:",n2.fivetimes())

}