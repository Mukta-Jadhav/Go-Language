package main
import ("fmt")
func main(){
	var n int=10;
	a,b:=0,1
	fmt.Println("Fibonacci sequence (%d terms):",n)
	for i:=0;i<n;i++{
		fmt.Printf("%d \t",a)
		a,b=b,a+b
	}

}
