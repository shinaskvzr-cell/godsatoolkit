package main
import "fmt"

type Queue struct{
	arr []int
	front int
	rear int
	size int
}

func NewQueue(cap int) *Queue{
	return &Queue{
		arr:make([]int,cap),
		front :0,
		rear:-1,
		size:0,
	}
}

func (q *Queue) Enqueue (val int){
	if q.size == len(q.arr)-1{
		fmt.Println("Queue Overflow")
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

func (q *Queue) Display(){
	if q.front == 0{
		fmt.Println("Queue is empty")
		return
	}
	for i:=q.front;i<=q.rear;i++{
		fmt.Printf("%d\n",q.arr[i])
	}
	fmt.Println()
}

func main(){

}