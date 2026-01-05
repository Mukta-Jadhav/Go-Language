package main
import ("fmt")
func main(){
	var a ,i,j int
	fmt.Print("enter a number:")
	fmt.Scan(&a)
    for i=0;i<=a;i++ {
		for j=0;j<=i;j++ {
			fmt.Println(i,"\n")
		}
	}
}