package maxArray
import "fmt"

func MaxInArray(){
	fmt.Println("\nAnalyze time complexity of a 'find max in array' function")
	arr := []int{98,12,33,67,89,34,666,7,-654}
	max := arr[0]
	for i:=0;i<len(arr);i++{
		if max < arr[i]{
			max = arr[i]
		}
	}
	fmt.Println(max)
}