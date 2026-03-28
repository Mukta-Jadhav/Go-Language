package main
import ("fmt")
func bubbleSort(arr []int) []int{
	for i:=0;i<=len(arr)-1;i++{
		for j:=0;j<=len(arr)-1-i;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
	}
}
return arr
}

func main(){
	var size int
	fmt.Print("enter the array size:")
	fmt.Scan(&size)
	numarray:=make([]int,size)
	fmt.Println("enter the array item:")
	for i:=0;i<size;i++{
		fmt.Scan(&numarray[i])
	}
	fmt.Println("sorted array is:",bubbleSort(numarray))

}