package main
import "fmt"

type Stack struct{
	arr []int
	top int
}

func NewStack (cap int) *Stack{
	return &Stack{
		arr : make([]int,cap),
		top :-1,
	}
}

func (s *Stack) Push(val int){
	if s.top == len(s.arr)-1{
		fmt.Println("Stack Overflow")
		return
	}
	s.top++
	s.arr[s.top]=val
}

func (s *Stack) Pop ()int{
	if s.top == -1{
		fmt.Println("Stack underflow")
		return -1
	}
	val := s.arr[s.top]
	s.top--
	return val
}

func ReverseArray(arr []int){
	s := NewStack(len(arr))
	for _,v := range arr{
		s.Push(v)
	}

	for i:=0;i<len(arr);i++{
		arr[i]=s.Pop()
	}
}

func main(){
	arr := []int{1,2,3,4,5,6,7,8,9}
	ReverseArray(arr)
	fmt.Println(arr)

}