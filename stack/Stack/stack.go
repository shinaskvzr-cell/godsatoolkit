package main
import "fmt"

type Stack struct{
	arr []int
	top int
}

func NewStack (cap int) *Stack{
	return &Stack{
		arr : make([]int,cap),
		top : -1,
	}
}

func (s *Stack) Push(val int){
	if s.top == len(s.arr){
		fmt.Println("Stack overflow")
		return
	}
	s.top++
	s.arr[s.top]=val
}

func (s *Stack) Pop()int{
	if s.top == -1{
		fmt.Println("Stack Underflow")
		return -1
	}
	val := s.arr[s.top]
	s.top--
	return val
}

func (s *Stack) Display(){
	if s.top == -1{
		fmt.Println("Stack is empty")
		return
	}
	for i:=s.top;i>=0;i--{
		fmt.Println(s.arr[i])
	}
}


func main(){
	s := NewStack(5)
	s.Push(10)
	s.Push(20)
	s.Push(40)
	s.Push(30)
	s.Display()

}
