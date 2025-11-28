package sliceMake
import "fmt"

func SliceMake(){
	fmt.Println("\nAllocate slices/maps using new and make")
	//make
	slice := make([]int,1,3)
	fmt.Println(slice)
	mapp := make(map[int]string)
	mapp[1]="Anshad"
	mapp[2]="Yadhu"
	mapp[3]="Shinas"
	fmt.Println(mapp)

	//new
	p := new(int)
	fmt.Println(p)
	*p = 999
	fmt.Println(*p)
}