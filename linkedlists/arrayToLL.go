package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (l *LinkedList) InsertAtFirst(val int){
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

func (l *LinkedList) FromArray(arr []int){
	for _,v := range arr{
		//l.InsertAtEnd(v)
		l.InsertAtFirst(v)
	}
}

func main(){
	list := LinkedList{}
	arr := []int{21,44,754,23,87,55,45,10,62}
	list.FromArray(arr)
	list.Display()

}