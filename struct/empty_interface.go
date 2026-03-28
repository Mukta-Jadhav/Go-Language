 
package main 
   
import "fmt" 
   
func myfun(a interface{}) { 
   
    // Extracting the value of a 
    val := a.(string) 
    fmt.Println("Value: ", val) 
} 
func main() { 
   
    var val interface { 
    } = "SPPU" 
       
    myfun(val) 
} 