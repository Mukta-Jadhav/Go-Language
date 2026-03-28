package main
import("fmt")
func max_min(array [] int)(int,int){
	var max,min int
	max=array[0]
	min=array[0]
	for _,a:=range array{
		if (max<a){
			max=a
		}
		if(min>a){
			min=a
		}
	}
	return max,min
}

func main(){
	var size int
	fmt.Print("Enter the array size:")
	fmt.Scan(&size)
	numarray:=make([] int,size)
	fmt.Println("Enter the array item:")
	for i:=0;i<size;i++{
		fmt.Scan(&numarray[i])
	}

	mx,mn:=max_min(numarray)
	fmt.Println("maximum of array item=",mx)
	fmt.Println("minimum of array item=",mn)
}
