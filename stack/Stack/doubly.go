package main
import "fmt"

type Node struct{
	data int
	next *Node
	prev *Node
}

type LinkedList struct{
	front *Node
	rear *Node
}

func (l *LinkedList) Push(val int){
	newNode := &Node{data:val}
	
	if l.front == nil{
		l.front = newNode
		l.rear = newNode
		return
	}

	l.rear.next = newNode
	newNode.prev = l.rear
	l.rear = newNode
}

func (l *LinkedList) Pop()int{
	if l.front == nil{
		fmt.Println("Queue Underflow")
		return -1
	}

	val := l.front.data
	l.front = l.front.next
	if l.front == nil {
		l.rear = nil
	} else {
		l.front.prev = nil
	}

	return val
}

func (l *LinkedList) Display(){
	if l.front == nil{
		fmt.Println("Queue is empty")
		return
	}
	current := l.front
	for current != nil{
		fmt.Printf("%d<->",current.data)
		current = current.next
	}
	fmt.Println("nil")
	
}


func main(){
	l := LinkedList{}
	l.Push(10)
	l.Push(20)
	l.Push(30)
	l.Push(50)
	l.Display()
	fmt.Println(l.Pop())
}