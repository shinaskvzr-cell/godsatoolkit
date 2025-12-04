package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type Queue struct{
	front *Node
}

func (q *Queue) Enqueue(val int){
	newNode := &Node{data:val}
	newNode.next = q.front
	q.front=newNode
}

func (q *Queue) Dequeue()int{
	if q.front == nil{
		fmt.Println("Queue underflow")
		return -1
	}
	val := q.front.data
	q.front = q.front.next
	return val
}

func (q *Queue) Display(){
	current := q.front
	if current == nil{
		fmt.Println("Queue is empty")
		return
	}
	for current != nil{
		fmt.Printf("%d->",current.data)
		current = current.next
	}
	fmt.Println("nil")
}


func main(){
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(90)
	q.Enqueue(20)
	q.Display()

}