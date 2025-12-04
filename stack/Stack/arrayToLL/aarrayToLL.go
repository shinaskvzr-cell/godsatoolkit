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
	newNode := &Node{data :val}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) Display(){
	current := l.head
	for current != nil{
		fmt.Printf("%d->",current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) FromArray(arr []int){
	for _,v := range arr{
		l.InsertAtFirst(v)
	}
}

func main(){
	l:= LinkedList{}
	arr := []int{1,2,3,4,5,6,67}
	l.FromArray(arr)
	l.Display()
}