package main

import (
	"fmt"

	"github.com/JosueMolinaMorales/algorithms/internal/data_structures/heap"
	"github.com/JosueMolinaMorales/algorithms/internal/sorting"
)

func main() {
	// Sorting algorithms
	arr := []int{10, 3, 4, 1, 5, 2, 6, 2, 4, 6}
	fmt.Println("Unsorted arr:", arr)
	sorting.MergeSort(&arr)
	fmt.Println("Sorted arr:", arr)

	arr = []int{10, 3, 4, 1, 5, 2, 6, 2, 4, 6}
	// Create Heap
	h := heap.Heapify(arr, heap.MaxHeap)
	fmt.Println("Max heap:", h.Heap)
	heap.PrintHeap(h.Heap)

	// Insert into heap
	h.Insert(7)
	fmt.Println("Max heap after insert:", h.Heap)
	heap.PrintHeap(h.Heap)

	// Pop from heap
	el, _ := h.Pop()
	fmt.Println("Max heap after pop:", h.Heap)
	fmt.Println("Popped element:", el)
	heap.PrintHeap(h.Heap)
}
