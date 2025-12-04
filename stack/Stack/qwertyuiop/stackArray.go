package main
import "fmt"

type Stack struct{
	arr []int
	top int
}

func NewArray (size int) *Stack{
	return &Stack{
		arr : make([]int,size),
		top :-1,
	}
}

func (s *Stack) Push (val int){
	if s.top == len(s.arr)-1{
		fmt.Println("Stack Overflow")
		return
	}
	s.top++
	s.arr[s.top]=val
}

func (s *Stack) Pop()int{
	if s.top == -1{
		fmt.Println("Stack underflow")
		return -1
	}
	val := s.arr[s.top]
	s.top--
	return val
}

func (s *Stack) Peek()int{
	if s.top == -1{
		fmt.Println("Stack underflow")
		return -1
	}
	return s.arr[s.top]
}

func (s *Stack) Display(){
	if s.top == -1{
		fmt.Println("Stack underflow")
		return
	}
	for i:=s.top;i>=0;i--{
		fmt.Println(s.arr[i])
	}
	fmt.Println("nil")
}

func main(){
	s := NewArray(5)
	s.Push(10)
	s.Push(70)
	s.Push(30)
	s.Push(99)
	s.Display()
}