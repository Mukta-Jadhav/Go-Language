package main
import("fmt")
func main(){
	var a,b int=5,2
	fmt.Println("befor swapping element are:\n")
	fmt.Println("a:%d,b:%d",a,b)
	a,b=b,a
	fmt.Println("after swapping element are:\n")
	fmt.Println("a:%d,b:%d",a,b)
}