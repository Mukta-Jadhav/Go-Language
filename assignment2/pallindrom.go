package main
import("fmt")

func pallindrome(a int) string{
     b:=a
	 rev:=0
	for a>0 {
     digit:=a%10
	 rev=(rev*10)+digit
	 a=a/10
	 }
	 //fmt.Printf("rev:%d",rev)
	 if b==rev{
		//fmt.Println("%d is pallindrome number",b)
		return fmt.Sprintf("%d is pallindrome number",b)
	 } else{
       // fmt.Println("%d is not pallindrome number",b)
	   return fmt.Sprintf("%d is not pallindrome number",b)
	 }
}

func main(){
  	var x int
	fmt.Print("enter a number:")
	fmt.Scan(&x)
	result:=pallindrome(x)
	fmt.Println(result)
}