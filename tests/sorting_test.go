package tests

import (
	"reflect"
	"testing"

	"github.com/JosueMolinaMorales/algorithms/internal/sorting"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order array",
			input:    []int{3, 1, 4, 2, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicates",
			input:    []int{3, 2, 1, 3, 2},
			expected: []int{1, 2, 2, 3, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := make([]int, len(test.input))
			copy(arr, test.input)
			sorting.QuickSort(&arr)
			if !reflect.DeepEqual(arr, test.expected) {
				t.Errorf("QuickSort(%v) = %v, expected %v", test.input, arr, test.expected)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order array",
			input:    []int{3, 1, 4, 2, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicates",
			input:    []int{3, 2, 1, 3, 2},
			expected: []int{1, 2, 2, 3, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := make([]int, len(test.input))
			copy(arr, test.input)
			sorting.MergeSort(&arr)
			if !reflect.DeepEqual(arr, test.expected) {
				t.Errorf("MergeSort(%v) = %v, expected %v", test.input, arr, test.expected)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order array",
			input:    []int{3, 1, 4, 2, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := make([]int, len(test.input))
			copy(arr, test.input)
			sorting.SelectionSort(&arr, len(arr))
			if !reflect.DeepEqual(arr, test.expected) {
				t.Errorf("SelectionSort(%v) = %v, expected %v", test.input, arr, test.expected)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order array",
			input:    []int{3, 1, 4, 2, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := make([]int, len(test.input))
			copy(arr, test.input)
			sorting.InsertionSort(&arr, len(arr))
			if !reflect.DeepEqual(arr, test.expected) {
				t.Errorf("InsertionSort(%v) = %v, expected %v", test.input, arr, test.expected)
			}
		})
	}
}
