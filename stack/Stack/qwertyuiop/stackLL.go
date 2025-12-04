package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type Stack struct{
	top *Node
}

func (s *Stack) Push (val int){
	newNode := &Node{data:val}
	newNode.next=s.top
	s.top=newNode
}

func (s *Stack) Pop() int{
	if s.top == nil{
		fmt.Println("Stack underflow")
		return -1
	}
	val := s.top.data
	s.top = s.top.next
	return val
}

func (s *Stack) Peek() int{
	if s.top == nil{
		fmt.Println("Stack is empty")
		return -1
	}
	return s.top.data
}

func (s *Stack) Display (){
	if s.top == nil{
		fmt.Println("Stack is empty")
		return
	}
	current := s.top
	for current != nil{
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println("Nil")
}


func main(){
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Display()

}