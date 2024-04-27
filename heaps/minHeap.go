package main
import "fmt"

type MinHeap struct {
	array []int
}

func (h *MinHeap) insert(key int){
	h.array = append(h.array, key)
	h.minHeapifyUp(len(h.array)-1)
}

func (h *MinHeap) Extract()int{
	extractedItem := h.array[0]
	l := len(h.array)-1

		//if array is empty
		if len(h.array) == 0 {
			fmt.Println("cannot extract because array length is zero")
			return -1
		}
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.minHeapifyDown(0)
	return extractedItem
}

func (h *MinHeap) minHeapifyUp(index int){
	for h.array[parent(index)] > h.array[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) minHeapifyDown(index int){
	lastIndex := len(h.array)-1
	l,r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex{
		if l == lastIndex{
			childToCompare = l
		}else if h.array[l] < h.array[r]{
			childToCompare = l
		}else{
			childToCompare = r
		}
		if h.array[index] > h.array[childToCompare]{
			h.swap(index, childToCompare)
			index = childToCompare
			l,r = left(index), right(index)
		}else {
			return
		}
	}
}
func parent(i int) int {
	return(i-1)/2
}
// Get the left child's index
func left(i int) int{
	return 2*i + 1
}
// Get the right child's index
func right(i int) int{
	return 2*i + 2
}

func (h *MinHeap) swap(i1, i2 int){
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main(){
 m := &MinHeap{}
buildHeap := []int{32,38,1,2,56,12,64}
 for _, v := range buildHeap {
	m.insert(v)
	fmt.Println(m)
 }
 for i := 0; i <= len(buildHeap) - 1 ; i++ {
	m.Extract()
	fmt.Println(m)
}
}