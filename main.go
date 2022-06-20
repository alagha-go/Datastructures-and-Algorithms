package main

import (
	"DSAL/lib"
	"fmt"
)


func main() {
	var heap lib.Heap
	slice := []int{10, 20, 30, 21, 43, 19, 28, 31, 67, 47, 86, 91}

	for _, value := range slice {
		heap.InsertMax(value)
		fmt.Println(heap.Slice)
	}

	for index:=0; index<10; index++ {
		heap.ExtractMax()
		fmt.Println(heap.Slice)
	}

	heap.ExtractMax()
	fmt.Println(heap.Slice)
}