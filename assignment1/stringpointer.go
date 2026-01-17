package main
import("fmt")
func main(){
	var str1,str2 string
	fmt.Println("Enter a string1:")
	fmt.Scan(&str1)
	fmt.Println("Enter a string 2:")
	fmt.Scan(&str2)
	p1:=&str1
	p2:=&str2
	concate:=*p1+"\t"+*p2
	fmt.Printf("concate string is:%s",concate)
}