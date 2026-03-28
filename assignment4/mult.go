package main

import ("fmt")

type number struct{}


func (n number) multiply(x int ,y int) int{
	return x *y
}

func main(){
	num:=number{}
	result:=num.multiply(10,5)
	fmt.Printf("multiplication of %d and %d is:%d",10,5,result)

}
