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

func (l *LinkedList) InsertAtEnd (val int){
	newNode := &Node{data :val}
	if l.head == nil{
		l.head=newNode
		return
	}
	current := l.head
	for current.next != nil{
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Display (){
	current := l.head
	for current != nil{
		fmt.Printf("%d ->",current.data)
		current=current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) Delete(val int){
	if l.head == nil{
		return
	}

	if l.head.data == val{
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil && current.next.data != val{
		current=current.next
	}

	if current.next == nil{
		return
	}

	current.next = current.next.next
}

func (l *LinkedList) Search(val int) bool {
	current := l.head

	for current != nil {
		if current.data == val {
			return true
		}
		current = current.next
	}

	return false // not found
}


func main(){
	l := &LinkedList{}
	l.InsertAtFirst(10)
	l.InsertAtFirst(20)
	l.InsertAtFirst(30)
	l.Display()
	l.InsertAtEnd(99)
	l.Display()
	l.Delete(20)
	l.Display()
	fmt.Println(l.Search(99))
	l.Display()

}