package heap

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

func (h *Heap) buildMaxHeap() {
	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *Heap) buildMinHeap() {
	for i := h.size/2 - 1; i >= 0; i-- {
		h.minHeapify(i)
	}
}

func (h *Heap) minHeapify(i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

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

func (h *Heap) heapifyDown(i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

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
		h.heapifyDown(largest)
	}
}

func parent(index int) int {
	return (index - 1) / 2
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
