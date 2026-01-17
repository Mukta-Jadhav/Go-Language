package main
import ("fmt")
func main(){
	var n ,a int
	fmt.Printf("enter a number of rows:")
	fmt.Scanf("%d",&n)
    for row:=1;row<=n;row++ {
		a=1
		for i:=1;i<=40-row;i++ {
			fmt.Printf(" ")
			
		}
		for i:=1;i<=row;i++{
           fmt.Printf("%d",a)
		   a=a*(row-i)/i
		}
		
		fmt.Print("\n")
	}
}