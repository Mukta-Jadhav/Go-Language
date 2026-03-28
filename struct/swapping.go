package main
import("fmt")
type Number struct{
	a,b int
}

func (r *Number) swap_pr(){
	r.b=r.a+r.b
	r.a=r.b-r.a
	r.b=r.b-r.a
}

func (r Number) swap1(){
	r.b=r.a+r.b
	r.a=r.b-r.a
	r.b=r.b-r.a
}

func main(){
	n1:=Number{a:10,b:20}
	fmt.Printf("before swapping number are %d and %d\n",n1.a,n1.b)
	n1.swap1()
	fmt.Printf("after swapping number are %d and %d\n",n1.a,n1.b)
	fmt.Printf("before swapping number are %d and %d\n",n1.a,n1.b)
	n1.swap_pr()
	fmt.Printf("after swapping number are %d and %d\n",n1.a,n1.b)
}