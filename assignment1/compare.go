package main
import("fmt"
"strings"
)
func main(){
	var str1 string="golang"
	var str2 string="GOLANG"
	fmt.Println("string compare:",str1==str2)

	//using  compare ()
	fmt.Println(strings.Compare("TYBCA","TYBCA"))
}