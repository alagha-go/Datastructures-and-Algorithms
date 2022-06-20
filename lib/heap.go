package lib


type Heap struct {
	Slice []int
}

/// insert value to the heap
func (h *Heap) Insert(value int) {
	h.Slice = append(h.Slice, value)
	h.HipifyUp(len(h.Slice)-1)
}

/// heapify heap from bottom to top
func (h *Heap) HipifyUp(index int) {
	for h.Slice[Parent(index)] < h.Slice[index] {
		h.Swap(Parent(index), index)
		index = Parent(index)
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