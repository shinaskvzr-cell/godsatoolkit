package modifyVariable
import "fmt"

func Modify(x *int){
	*x = 99
}


func ModifyVariable(){
	fmt.Println("\nModify a variable using pointer")
	x := 10
	Modify(&x)
	fmt.Println(x)
}