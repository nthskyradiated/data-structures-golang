package main
import "fmt"

type Stack struct {
	items []int 
}

//push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//pop
func (s *Stack) Pop() int {
	l := len(s.items)-1
	removedItm := s.items[l]
	s.items = s.items[:l]
	return removedItm
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

}