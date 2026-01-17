package main
import ("fmt")

func swap(a ,b *int){
	var temp int

	temp=*a
	*a=*b 
	*b=temp

	}

	func main(){
		var a int=45
		var b int=35

		fmt.Printf("before swapping ,value of a is:%d and value of b is:%d \n",a,b)

		swap(&a,&b)
		fmt.Printf("before swapping ,value of a is:%d and value of b is:%d \n",a,b)

	}