package main
import("fmt")
func add(a,b int) int{
	result:=a+b
	return result

}

func main(){
	var x,y int
	fmt.Print("enter value of x:")
	fmt.Scan(&x)
	fmt.Print("enter value of y:")
	fmt.Scan(&y)
    addition:=add(x,y)
	fmt.Printf("addition of a:%d and b:%d is:%d",x,y,addition)
}