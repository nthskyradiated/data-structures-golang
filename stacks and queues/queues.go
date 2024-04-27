package main

import "fmt"


type Queue struct {
	items []int
}

//enqueue
func (q *Queue) enq(i int) {
	q.items = append(q.items, i)
}
//dequeue
func (q *Queue) dq() int {
removedItm := q.items[0]
q.items = q.items[1:]
return removedItm
}


func main(){
	myQ := Queue{}
	fmt.Println(myQ)
	myQ.enq(100)
	myQ.enq(200)
	myQ.enq(300)
	fmt.Println(myQ)
	myQ.dq()
	fmt.Println(myQ)
	myQ.dq()
	fmt.Println(myQ)
}