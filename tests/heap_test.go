package tests

import (
	"reflect"
	"testing"

	"github.com/JosueMolinaMorales/algorithms/internal/data_structures/heap"
)

func TestInsertOperation(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		insert   int
		heapType int
		expected []int
	}{
		{
			name:     "MaxHeap: Empty array",
			arr:      []int{},
			heapType: heap.MaxHeap,
			expected: []int{6},
			insert:   6,
		},
		{
			name:     "MaxHeap: Already max heap",
			arr:      []int{5, 4, 3, 2, 1},
			heapType: heap.MaxHeap,
			expected: []int{6, 4, 5, 2, 1, 3},
			insert:   6,
		},
		{
			name:     "MaxHeap: Random order",
			arr:      []int{3, 2, 1, 5, 4},
			heapType: heap.MaxHeap,
			expected: []int{6, 4, 5, 2, 3, 1},
			insert:   6,
		},
		{
			name:     "MinHeap: Empty array",
			arr:      []int{},
			heapType: heap.MinHeap,
			expected: []int{6},
			insert:   6,
		},
		{
			name:     "MinHeap: Already min heap",
			arr:      []int{1, 2, 3, 4, 5},
			heapType: heap.MinHeap,
			expected: []int{0, 2, 1, 4, 5, 3},
			insert:   0,
		},
		{
			name:     "MinHeap: Random order",
			arr:      []int{5, 4, 3, 2, 1},
			heapType: heap.MinHeap,
			expected: []int{0, 2, 1, 5, 4, 3},
			insert:   0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := heap.Heapify(test.arr, test.heapType)
			h.Insert(test.insert)
			if !reflect.DeepEqual(h.Heap, test.expected) {
				t.Errorf("Insert(%d) = %v, expected %v", test.insert, h.Heap, test.expected)
			}
		})
	}
}

func TestHeapify(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		heapType int
		expected []int
	}{
		{
			name:     "MaxHeap: Empty array",
			arr:      []int{},
			heapType: heap.MaxHeap,
			expected: []int{},
		},
		{
			name:     "MaxHeap: Already max heap",
			arr:      []int{5, 4, 3, 2, 1},
			heapType: heap.MaxHeap,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "MaxHeap: Random order",
			arr:      []int{3, 2, 1, 5, 4},
			heapType: heap.MaxHeap,
			expected: []int{5, 4, 1, 2, 3},
		},
		{
			name:     "MinHeap: Empty array",
			arr:      []int{},
			heapType: heap.MinHeap,
			expected: []int{},
		},
		{
			name:     "MinHeap: Already min heap",
			arr:      []int{1, 2, 3, 4, 5},
			heapType: heap.MinHeap,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "MinHeap: Random order",
			arr:      []int{5, 4, 3, 2, 1},
			heapType: heap.MinHeap,
			expected: []int{1, 2, 3, 5, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := heap.Heapify(test.arr, test.heapType)
			if !reflect.DeepEqual(h.Heap, test.expected) {
				t.Errorf("Heapify(%v) with heapType %d = %v, expected %v", test.arr, test.heapType, h.Heap, test.expected)
			}
		})
	}
}
