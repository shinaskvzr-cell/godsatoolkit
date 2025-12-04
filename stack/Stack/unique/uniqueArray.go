package main
import "fmt"

func main(){
	arr := []int{1,2,3,4,4,54,5,5,3,2,4,57,1,34,66,4,2,1,27,3765,4,45,674,8,81,9,7,6,5,42,32,3,1,2,2}
	mapp := make(map[int]bool)
	result := []int{}
	for _,v := range arr{
		if !mapp[v]{
			result = append(result,v)
			mapp[v]=true
		}
	}
	fmt.Println(result)
}