package main
import("fmt")
func main(){
	var a int
	var i int
	fmt.Println("Enter a number:")
	fmt.Scanf("%d",&a)
	fmt.Println("table of ",a, "is :")
	for i=1;i<=10;i++ {
		fmt.Println(a,"X",i,"=",a*i)
	}
}