package main
import "fmt"

type Queue struct{
	arr []int
}

func (q *Queue) Enqueue(val int){
	q.arr = append(q.arr,val)
}
 

func (q *Queue) Dequeue()int{
	if len(q.arr) == 0{
		fmt.Println("Queue underflow")
		return -1
	}
	val := q.arr[0]
	q.arr = q.arr[1:]
	return val
}

func (q *Queue) Peek()int{
	if len(q.arr) == 0{
		fmt.Println("Queue is empty")
		return -1
	}
	return q.arr[0]
}

func (q *Queue) Display(){
	if len(q.arr) == 0{
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println(q.arr)
}
func main(){
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(9870)
	q.Enqueue(450)
	q.Enqueue(90)
	q.Display()
}