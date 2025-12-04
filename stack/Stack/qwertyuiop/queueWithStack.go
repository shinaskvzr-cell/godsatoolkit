package main
import "fmt"

type Queue struct{
	arr []int
	front int
	rear int
	size int
}

func NewArray (cap int) *Queue {
	return &Queue{
		arr :make ([]int,cap),
		front :0,
		rear :-1,
		size :0,
	}
}

func (q *Queue) Enqueue(val int){
	if q.size == len(q.arr)-1{
		fmt.Println("Queue overflow")
		return
	}
	q.rear++
	q.arr[q.rear]=val
	q.size++
}

func (q *Queue) Dequeue()int{
	if q.size == 0{
		fmt.Println("Queue underflow")
		return -1
	}
	val := q.arr[q.front]
	q.front++
	q.size--
	return val
}

func (q *Queue) Peek()int{
	if q.size == 0{
		fmt.Println("Queue is empty")
		return -1
	}
	return q.arr[q.front]
}

func (q *Queue) Display(){
	if q.size == 0{
		fmt.Println("Queue is empty")
		return
	}
	for i:=q.front;i<=q.rear;i++{
		fmt.printf("%d",q.arr[i])
	}
	fmt.Println()
}

func main(){

}