package heap

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	// MaxHeap represents a max heap
	MaxHeap = iota
	// MinHeap represents a min heap
	MinHeap
)

// Heap is a data structure that represents a binary tree
type Heap struct {
	Heap     []int
	size     int
	heapType int
}

// NewHeap Creates a new heap with a given type
func NewHeap(heapType int) *Heap {
	return &Heap{heapType: heapType}
}

// Heapify creates a heap from an array
// Time complexity: O(n)
// Auxiliary space: O(1)
func Heapify(arr []int, heapType int) *Heap {
	h := &Heap{}
	h.Heap = arr
	h.size = len(arr)

	switch heapType {
	case MaxHeap:
		h.heapType = MaxHeap
		h.buildMaxHeap()
	case MinHeap:
		h.heapType = MinHeap
		h.buildMinHeap()
	default:
		panic("Invalid heap type")
	}

	return h
}

// Insert inserts a new element into the heap
func (h *Heap) Insert(element int) {
	// Add the element to the heap
	h.Heap = append(h.Heap, element)
	h.size++

	// Heapify to fix the heap
	i := h.size - 1
	for i > 0 && (h.heapType == MaxHeap && h.Heap[parent(i)] < h.Heap[i]) || (h.heapType == MinHeap && h.Heap[parent(i)] > h.Heap[i]) {
		h.Heap[i], h.Heap[parent(i)] = h.Heap[parent(i)], h.Heap[i]
		i = parent(i)
	}
}

// Peek returns the root of the heap
func (h *Heap) Peek() int {
	return h.Heap[0]
}

// Pop returns the top element of the heap, removing it from the heap
func (h *Heap) Pop() (int, error) {
	if h.size == 0 {
		return -1, errors.New("Heap is empty")
	}
	h.size--
	el := h.Heap[0]
	h.Heap = h.Heap[1:]
	switch h.heapType {
	case MaxHeap:
		h.heapType = MaxHeap
		h.buildMaxHeap()
	case MinHeap:
		h.heapType = MinHeap
		h.buildMinHeap()
	}
	return el, nil
}

// PrintHeap prints the heap in a tree-like structure
func PrintHeap(arr []int) {
	n := len(arr)
	if n == 0 {
		fmt.Println("Heap is empty")
		return
	}

	// Calculate the height of the binary tree
	height := int(math.Ceil(math.Log2(float64(n + 1))))

	// Calculate the maximum number of nodes in the last level
	maxNodesLastLevel := int(math.Pow(2, float64(height-1)))

	// Calculate the maximum number of characters for each element
	maxElemWidth := len(fmt.Sprint(arr[0]))
	for i := 1; i < n; i++ {
		elemWidth := len(fmt.Sprint(arr[i]))
		if elemWidth > maxElemWidth {
			maxElemWidth = elemWidth
		}
	}

	// Calculate the total width for each level
	totalWidth := maxElemWidth * maxNodesLastLevel

	// Print the heap in a tree-like structure
	level := 1
	nodesPrinted := 0

	for i := 0; i < n; i++ {
		if nodesPrinted == 0 {
			fmt.Println()
			padding := totalWidth / (int(math.Pow(2, float64(level))) + 1)
			fmt.Print(strings.Repeat(" ", padding))
		}

		fmt.Printf("%*d", maxElemWidth, arr[i])

		nodesPrinted++

		if nodesPrinted == int(math.Pow(2, float64(level-1))) {
			level++
			nodesPrinted = 0
		} else {
			padding := totalWidth / (int(math.Pow(2, float64(level))) + 1)
			fmt.Print(strings.Repeat(" ", padding))
		}
	}

	fmt.Println()
}

func (h *Heap) buildMaxHeap() {
	for i := h.size/2 - 1; i >= 0; i-- {
		h.maxHeapify(i)
	}
}

func (h *Heap) buildMinHeap() {
	for i := h.size/2 - 1; i >= 0; i-- {
		h.minHeapify(i)
	}
}

func (h *Heap) minHeapify(i int) {
	smallest := i
	left := left(i)
	right := right(i)

	// Compare the current node with its left child
	if left < h.size && h.Heap[left] < h.Heap[smallest] {
		smallest = left
	}

	// Compare the current node with its right child
	if right < h.size && h.Heap[right] < h.Heap[smallest] {
		smallest = right
	}

	if smallest != i {
		h.Heap[i], h.Heap[smallest] = h.Heap[smallest], h.Heap[i]
		h.minHeapify(smallest)
	}
}

func (h *Heap) maxHeapify(i int) {
	largest := i
	left := left(i)
	right := right(i)

	// Compare the current node with its left child
	if left < h.size && h.Heap[left] > h.Heap[largest] {
		largest = left
	}

	// Compare the current node with its right child
	if right < h.size && h.Heap[right] > h.Heap[largest] {
		largest = right
	}

	// If the largest element is not the current node, swap them
	if largest != i {
		h.Heap[i], h.Heap[largest] = h.Heap[largest], h.Heap[i]
		h.maxHeapify(largest)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}
