package main
import ("fmt")
func main(){
	var num int
	fmt.Println("enter a number:")
	fmt.Scanf("%d",&num)
	if (num<10 && num>-10){
		fmt.Printf("%d is single digit",num)
	} else{
		fmt.Printf("%d is  not single digit",num)
	} 
}