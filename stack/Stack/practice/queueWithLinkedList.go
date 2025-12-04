package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type Queue struct{
	front *Node
	rear *Node
}

func (q *Queue) Enqueue(val int){
	newNode := &Node{data:val}
	if q.front == nil{
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue ()int{
	if q.front == nil{
		fmt.Println("Queue underflow")
		return -1
	}
	val := q.front.data
	q.front = q.front.next

	if q.front == nil{
		q.rear = nil
	}

	return val
}

func (q *Queue) Peek()int{
	if q.front == nil{
		fmt.Println("Queue is empty")
		return -1
	}
	return q.front.data
}

func (q *Queue) Display(){
	if q.front == nil{
		fmt.Println("Queue is empty")
		return
	}
	current := q.front
	for current != nil{
		fmt.Printf("%d->",current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main(){
	q:= Queue{}
	q.Enqueue(10)
	q.Enqueue(120)
	q.Enqueue(650)
	q.Display()

}