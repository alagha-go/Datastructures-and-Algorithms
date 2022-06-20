package lib

import "fmt"


type Heap struct {
	Slice []int
}

/// insert value to the heap
func (h *Heap) InsertMax(value int) {
	h.Slice = append(h.Slice, value)
	h.MaxHipifyUp(len(h.Slice)-1)
}

/// heapify heap from bottom to top
func (h *Heap) MaxHipifyUp(index int) {
	for h.Slice[Parent(index)] < h.Slice[index] {
		h.Swap(Parent(index), index)
		index = Parent(index)
	}
}

/// extracts largest value from the heap slice
func (h *Heap) ExtractMax() int {
	if len(h.Slice) == 0 {
		fmt.Println("cannot extract from an empty slice")
		return -1
	}
	value := h.Slice[0]
	lastindex := len(h.Slice)-1
	h.Slice[0] = h.Slice[lastindex]
	h.Slice = h.Slice[:lastindex]
	return value
}

// heapify from top to bottom
func (h *Heap) MaxHipifyDown(index int) {
	lastindex := len(h.Slice)-1
	left, right := LeftChild(index), RightChild(index)
	child := 0

	// loop if the current index has at least one child
	for left <= lastindex {
		// when the left child is the only child
		if left == lastindex {
			child = left
		}else if h.Slice[left] > h.Slice[right] {
			// when the left child is larger than the right child
			child = left
		}else {
			// when the right child is larger than the left child
			child = right
		}

		// compare the current index to larger child and swap
		if h.Slice[index] < h.Slice[child] {
			h.Swap(index, child)
			index = child
			left, right = LeftChild(index), RightChild(index)
		}else {
			return
		}
	}
}


/// get parent index
func Parent(index int) int {
	return (index-1) / 2
}


/// get the left child
func LeftChild(index int) int {
	return index*2+1
}


/// get the right child
func RightChild(index int) int {
	return index*2+2
}

/// swap values of the slice
func (h *Heap) Swap(index1, index2 int) {
	h.Slice[index1], h.Slice[index2] = h.Slice[index2], h.Slice[index1]
}