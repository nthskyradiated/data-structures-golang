package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int){
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}


// Extract returns the largest key and removes it from the heap

func (h *MaxHeap) Extract() int {
	extractedItem := h.array[0]
	l := len(h.array)-1

	//if array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is zero")
		return -1
	}

	//take the last indexed item out and put it in the root 
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)
	return extractedItem
}

// maxHeapifyUp will heapify from the bottom up 

func (h *MaxHeap) maxHeapifyUp(index int){
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapify down will heapit from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l,r := left(index), right(index)
	childToCompare := 0

	// Loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex{		 // When there's only only child
		childToCompare = l
			} else if h.array[l] > h.array[r] {		// When left child is larger
				childToCompare = l
			} else {
			childToCompare = r 		// When right child is larger
			}
			// Compare array value of current index to larger child and swap if smaller
			if h.array[index] < h.array[childToCompare]{
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
			} else {
				return
			}
	}
}


// Get parent index //left child is always an odd number while the right child is an even number
func parent(i int) int {
	return (i-1)/2
}

// Get the left child's index
func left(i int) int{
	return 2*i + 1
}
// Get the right child's index
func right(i int) int{
	return 2*i + 2
}

// Swap keys in the array
func(h *MaxHeap) swap (i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func main(){
	m := &MaxHeap{}
	buildHeap := []int{51,20,30,40,50,28,104,75,67,1111}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i <= len(buildHeap) - 1 ; i++ {
		m.Extract()
		fmt.Println(m)
	}

}