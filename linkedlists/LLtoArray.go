package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (l *LinkedList) InsertAtFirst (val int){
	newNode := &Node{data:val}
	newNode.next=l.head
	l.head=newNode
}

func (l *LinkedList) InsertAtEnd (val int){
	newNode := &Node{data:val}
	if l.head == nil{
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil{
		current=current.next
	}
	current.next = newNode
}

func (l *LinkedList) Display(){
	current := l.head
	if current == nil{
		return
	}
	for current != nil{
		fmt.Printf("%d ->",current.data)
		current=current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) LLtoArray() []int{
	result := []int{}
	current := l.head
	for current != nil{
		result = append(result,current.data)
		current=current.next
	}
	return result
}

func main (){
	l := LinkedList{}
	l.InsertAtFirst(10)
	l.InsertAtFirst(30)
	l.InsertAtEnd(1002)
	l.InsertAtFirst(40)
	l.Display()
	fmt.Println(l.LLtoArray())
}