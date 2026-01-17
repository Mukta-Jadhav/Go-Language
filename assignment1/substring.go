package main
import("fmt"
      "strings")
func main(){
	var str1,str2 string
	fmt.Println("Enter a string1:")
	fmt.Scan(&str1)
	fmt.Println("Enter a string 2:")
	fmt.Scan(&str2)
	if strings.Contains(str1,str2){
		fmt.Println("String2 is a substring of string1")
	}else{
		fmt.Println("String2 is a substring of string1")
		
	}
}
