package main
import("fmt")

func multivalue(a int ,s string)(int,string){
	var sum int
	var msg string
	
	
	 sum=a+50
	
	 msg=s+"\t"+"welcome to BJS"
	 return sum,msg
	   
 
}

func main(){
	var b int
	fmt.Println("enter number:")
	fmt.Scan(&b)

	fmt.Println(multivalue(b,"hi"))

}