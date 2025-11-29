package main
import "fmt"

type DNode struct{
	data int
	next *DNode
	prev *DNode
}

type DLinkedList struct{
	head *DNode
}

func (l *DLinkedList) InsertAtFirst(val int){
	newNode := &Node{data :val}
	if l.head == nil{
		l.head = newNode
	}
	newNode.next=l.head
	l.head.prev=newNode
	l.head=newNode
}


func (l *DLinkedList) DeleteLast(){
	if l.head == nil{
		return
	}
	if l.head.next == nil{
		l.head=nil
		return
	}
	current := l.head
	for current != nil{
		current=current.next
	}
	current.prev.next=nil
}

func (l *DLinkedList) Display(){
	current := l.head
	for current != nil{
		fmt.Printf("%d <-> ",current.data)
		current=current.next
	}
	fmt.Println("nil")

}

func main(){
	d := DLinkedList{}
	d.InsertAtFirst(10)
	d.InsertAtFirst(70)
	d.InsertAtFirst(20)
	d.InsertAtFirst(50)
	d.InsertAtFirst(80)
	d.Display()

}