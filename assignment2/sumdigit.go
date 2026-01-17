package main
import("fmt")
func sumdigits(num int)int{
remainder:=0
sum:=0
for num != 0{
	remainder=num%10
	sum=sum+remainder
	num=num/10

	}
return sum
}

func main(){
	var num int
	fmt.Println("enter a number:")
	fmt.Scan(&num)
	fmt.Printf("sum of digit of %d is:%d",num,sumdigits(num))
}