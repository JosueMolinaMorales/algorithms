package main

import (
	"fmt"

	"github.com/JosueMolinaMorales/algorithms/internal/sorting"
)

func main() {
	// Sorting algorithms
	arr := []int{10, 3, 4, 1, 5, 2, 6, 2, 4, 6}
	fmt.Println("Unsorted arr:", arr)
	sorting.MergeSort(&arr)
	fmt.Println("Sorted arr:", arr)
}
